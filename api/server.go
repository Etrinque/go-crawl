package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/etrinque/go-crawl/util"
	"github.com/gorilla/websocket"
)

type CrawlRequest struct {
	URL      string `json:"url"`
	Workers  int    `json:"numWorkers"`
	MaxPages int    `json:"numPages"`
}
type CrawlResponse struct {
	Pages      map[string]int  `json:"pages"`
	TotalPages int             `json:"totalPages"`
	Duration   string          `json:"duration"`
	LogSummary util.LogSummary `json:"logSummary"`
	Error      string          `json:"error,omitempty"`
}
type ProgressUpdate struct {
	CurrentPage int            `json:"currentPage"`
	MaxPages    int            `json:"maxPages"`
	PagesFound  map[string]int `json:"pagesFound"`
}

var upgrader = websocket.Upgrader{}

func startServer() {

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/api/crawl", handleCrawl)
	mux.HandleFunc("/api/logs", handleLogs)

	// gracefull shutdown
	go func() {
		fmt.Println("Starting server on port http://localhost:8080")
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server Failue: %+v", err)
		}
	}()
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	<-stopChan
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

}

func handleHome(w http.ResponseWriter, req *http.Request) {
	// serve gui
}

func handleCrawl(w http.ResponseWriter, req *http.Request) {
	// send parameters to crawlerConfig
}

func handleLogs(w http.ResponseWriter, req *http.Request) {
	// send logs to logOutput
}
