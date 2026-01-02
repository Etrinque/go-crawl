package util

import (
	"fmt"
	"sync"
	"time"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

func (l LogLevel) String() string {
	switch l {
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarning:
		return "WARNING"
	case LogLevelError:
		return "ERROR"
	case LogLevelDebug:
		return "DEBUG"
	default:
		return "Unknown"
	}
}

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     LogLevel  `json:"level"`
	Msg       string    `json:"message"`
	Err       string    `json:"error,omitempty"`
	Url       string    `json:"url,omitempty"`
}

type Logger struct {
	entries []LogEntry
	mu      sync.RWMutex
	logChan chan LogEntry
}

// NewLogger creates a new Logger with unbounded entries slice and buffered channel;
func NewLogger() *Logger {
	return &Logger{
		entries: make([]LogEntry, 0),
		logChan: make(chan LogEntry, 512),
	}
}

func (l *Logger) Log(level LogLevel, msg string, url string, err error) {
	if level < LogLevelDebug || level > LogLevelError {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Msg:       msg,
		Url:       url,
	}
	if err != nil {
		entry.Err = err.Error()
	}

	l.mu.Lock()
	l.entries = append(l.entries, entry)
	l.mu.Unlock()

	// send over channel for GUI to display
	// NOTE: this is a blocking call
	select {
	case l.logChan <- entry:
	default:
	}

	l.printToConsole(entry)
}

func (l *Logger) printToConsole(entry LogEntry) {
	// ANSI color codes
	const (
		colorReset  = "\033[0m"
		colorRed    = "\033[31m"
		colorYellow = "\033[33m"
		colorBlue   = "\033[34m"
		colorGray   = "\033[90m"
	)

	var color string
	switch entry.Level {
	case LogLevelDebug:
		color = colorGray
	case LogLevelInfo:
		color = colorBlue
	case LogLevelWarning:
		color = colorYellow
	case LogLevelError:
		color = colorRed
	default:
		color = colorReset
	}

	timestamp := entry.Timestamp.Format("15:04:05")

	if entry.Url != "" && entry.Err != "" {
		fmt.Printf("%s[%s] %s%s - %s | URL: %s | Error: %s\n",
			color, timestamp, entry.Level, colorReset, entry.Msg, entry.Url, entry.Err)
	} else if entry.Url != "" {
		fmt.Printf("%s[%s] %s%s - %s | URL: %s\n",
			color, timestamp, entry.Level, colorReset, entry.Msg, entry.Url)
	} else if entry.Err != "" {
		fmt.Printf("%s[%s] %s%s - %s | Error: %s\n",
			color, timestamp, entry.Level, colorReset, entry.Msg, entry.Err)
	} else {
		fmt.Printf("%s[%s] %s%s - %s\n",
			color, timestamp, entry.Level, colorReset, entry.Msg)
	}
}

func (l *Logger) Debug(message, url string) {
	l.Log(LogLevelDebug, message, url, nil)
}
func (l *Logger) Info(message, url string) {
	l.Log(LogLevelInfo, message, url, nil)
}
func (l *Logger) Warning(message, url string, err error) {
	l.Log(LogLevelWarning, message, url, err)
}
func (l *Logger) Error(message, url string, err error) {
	l.Log(LogLevelError, message, url, err)
}

func (l *Logger) EntriesByLevel(level LogLevel) []LogEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()

	filtered := make([]LogEntry, 0)
	for _, entry := range l.entries {
		if entry.Level == level {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

func (l *Logger) Entries() []LogEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()
	entries := make([]LogEntry, len(l.entries))
	copy(entries, l.entries)
	return entries
}

func (l *Logger) ErrorCount() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	count := 0
	for _, entry := range l.entries {
		if entry.Level == LogLevelError {
			count++
		}
	}
	return count
}

func (l *Logger) WarningCount() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	count := 0
	for _, entry := range l.entries {
		if entry.Level == LogLevelWarning {
			count++
		}
	}
	return count
}
func (l *Logger) InfoCount() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	count := 0
	for _, entry := range l.entries {
		if entry.Level == LogLevelInfo {
			count++
		}
	}
	return count
}

func (l *Logger) DebugCount() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	count := 0
	for _, entry := range l.entries {
		if entry.Level == LogLevelDebug {
			count++
		}
	}
	return count
}

func (l *Logger) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	// keep mem around for re-use
	l.entries = l.entries[:0]
}

func (l *Logger) LogChan() <-chan LogEntry {
	return l.logChan
}

func (l *Logger) Close() {
	close(l.logChan)
}

type LogSummary struct {
	TotalEntries int `json:"totalEntries"`
	ErrorCount   int `json:"errorCount"`
	WarningCount int `json:"warningCount"`
	InfoCount    int `json:"infoCount"`
	DebugCount   int `json:"debugCount"`
}

func (l *Logger) Summary() LogSummary {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return LogSummary{
		TotalEntries: len(l.entries),
		ErrorCount:   l.ErrorCount(),
		WarningCount: l.WarningCount(),
		InfoCount:    l.InfoCount(),
		DebugCount:   l.DebugCount(),
	}
}
