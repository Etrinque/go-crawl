package api

import (
	"fmt"
	"log"
	"net/http"

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
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/api/crawl", handleCrawl)
	http.HandleFunc("/api/logs", handleLogs)

	fmt.Println("Starting server on port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, req *http.Request) {
	// serve single page gui
}

func handleCrawl(w http.ResponseWriter, req *http.Request) {

}

func handleLogs(w http.ResponseWriter, req *http.Request) {}
