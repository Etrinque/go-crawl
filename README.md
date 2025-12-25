# Go-Crawl

<img src="https://github.com/Etrinque/go-crawl/blob/main/assets/logo.svg" width="100" height="100" alt="Go-Crawl Logo"/> 


A concurrent, thread-safe web crawler built in Go that demonstrates efficient URL extraction and crawling patterns.

## Overview

Go-Crawl is a command-line web crawler that recursively discovers and maps internal links on a website. It uses
goroutines and channels to crawl multiple pages concurrently while maintaining thread safety through proper
synchronization primitives.

## Features

- **Concurrent Crawling**: Configurable worker pool using goroutines and buffered channels
- **Thread-Safe State Management**: Uses `sync.Mutex` to safely track visited pages across goroutines
- **Domain Restriction**: Only crawls pages within the same domain as the starting URL
- **URL Normalization**: Handles different URL formats (http/https, trailing slashes, case sensitivity)
- **Link Analysis**: Generates a report showing how many times each page was referenced
- **Graceful Error Handling**: Continues crawling even when individual pages fail
- **Custom Sorting**: Implements merge sort to rank pages by reference count

## Technical Highlights

This project demonstrates:

- **Concurrency Patterns**: Worker pool pattern with a semaphore-like channel for controlling concurrent HTTP requests
- **Synchronization**: Proper use of `sync.Mutex` and `sync.WaitGroup` for coordinating goroutines
- **HTML Parsing**: Uses `golang.org/x/net/html` for robust HTML tokenization
- **URL Handling**: Leverages `net/url` for parsing and resolving relative URLs
- **Algorithm Implementation**: Custom merge sort (O(n log n)) for sorting results

## Installation

### Prerequisites

- Go 1.22 or later

### Build from Source

```bash
# Clone the repository
git clone https://github.com/etrinque/go-crawl
cd go-crawl

# Build the binary
go build -o crawl

# Or install to $GOPATH/bin
go install
```

## Usage

```bash
./crawl <url> <numWorkers> <maxPages>
```

### Arguments

- `url`: The starting URL to crawl (must include scheme: http:// or https://)
- `numWorkers`: Number of concurrent workers (recommended: 5–20)
- `maxPages`: Maximum number of pages to crawl (prevents infinite crawling)

### Examples

```bash
# Crawl blog.boot.dev with 10 workers, max 100 pages
./crawl https://blog.boot.dev 10 100

# Crawl example.com with 5 workers, max 50 pages
./crawl https://example.com 5 50
```

### Sample Output

```
Starting crawl of https://blog.boot.dev
Workers: 10 | Max pages: 100

========================================
  CRAWL REPORT: https://blog.boot.dev
========================================

Found 47 internal link(s) to blog.boot.dev
Found 23 internal link(s) to blog.boot.dev/path/to/article
Found 19 internal link(s) to blog.boot.dev/about
Found 12 internal link(s) to blog.boot.dev/contact
...

Total pages crawled: 100
```

## Architecture

### Core Components

**`config`**: Central state manager holding:

- Visited page map (protected by mutex)
- Concurrency control channel (semaphore pattern)
- WaitGroup for goroutine coordination

**`crawl()`**: Recursive crawler that:

1. Acquires worker slot from a channel
2. Fetches and parses HTML
3. Extracts URLs
4. Spawns goroutines for discovered links

**`getHTML()`**: HTTP client with:

- Timeout configuration
- Content-type validation
- Error status handling

**`getURLsFromHTML()`**: HTML parser that:

- Traverses the DOM tree recursively
- Extracts href attributes
- Resolves relative URLs

## Design Decisions

### Why Merge Sort?

While Go's standard library provides `sort.Slice()`, implementing merge sort demonstrates:

- Understanding of O(n log n) sorting algorithms
- Ability to implement classic CS algorithms
- Custom comparison logic for domain-specific types

### Why Buffered Channel for Concurrency Control?

Using a buffered channel as a semaphore provides:

- Simple, idiomatic concurrency limiting
- No additional dependencies
- Built-in blocking behavior

### Why Track Reference Counts?

Counting how many times each page is referenced helps identify:

- Most important pages (high reference count = hub pages)
- Site structure and navigation patterns
- Potential dead links (reference count = 1)

## Performance Considerations

- **Worker Pool Size**: Too many workers may overwhelm the target server or exhaust local resources. Recommended: 5–20
  workers
- **Max Pages**: Prevents runaway crawling. Consider the target site's size
- **HTTP Timeouts**: 10-second timeout prevents hanging on slow/unresponsive servers
- **Memory Usage**: Scales with number of unique pages (approximately 200-500 bytes per URL)

## Limitations

- Does not respect `robots.txt` (intentionally simplified for educational purposes)
- No rate limiting between requests to the same domain
- Does not handle JavaScript-rendered content
- Only follows `<a href>` links (ignores forms, redirects, etc.)
- No persistence - state is lost when the program exits

## Testing

Run tests with:

```bash
go test -v
```

Tests cover:

- URL normalization edge cases
- Merge sort correctness
- Error handling for invalid inputs

## Future Enhancements

Potential improvements for learning:

- Add `robots.txt` parsing and respect
- Implement rate limiting per domain
- Add command-line flags for configuration
- Export results to JSON/CSV
- Add HTTP caching headers support
- Visualize site structure as a graph

## License

MIT License ~

## Contributing

This is a portfolio project, but feedback and suggestions are welcome via issues.

Icon generated by Claude.ai.

---
