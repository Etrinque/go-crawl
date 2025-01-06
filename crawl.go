package main

import (
	"fmt"
	"net/url"
	"sync"
)

var errLog []error

type config struct {
	pages    map[string]int
	maxPages int
	root     *url.URL
	mut      *sync.Mutex
	wg       *sync.WaitGroup
	ch       chan struct{}
}

// NewConfig returns a new concurrency config. Worker pool size determined at initialization
// Channel buffered to worker-pool size.
func (c *config) NewConfig(root *url.URL, numWorkers int, maxPages int) *config {

	config := &config{
		pages:    make(map[string]int),
		maxPages: maxPages,
		root:     root,
		wg:       &sync.WaitGroup{},
		mut:      &sync.Mutex{},
		ch:       make(chan struct{}, numWorkers),
	}
	return config
}

// pagesLen Concurrent safe measure of page-map
func (c *config) pagesLen() int {
	c.mut.Lock()
	defer c.mut.Unlock()
	return len(c.pages)
}

// addVisited checks if the current page has been visited, false (not visited) by default.
// adds the page to the map. Concurrency Safe
func (c *config) addVisited(normCurUrl string) bool {
	c.mut.Lock()
	defer c.mut.Unlock()

	if _, ok := c.pages[normCurUrl]; ok {
		c.pages[normCurUrl]++
		return false
	}

	c.pages[normCurUrl] = 1
	return true
}

// Crawl is responsible for fetching -> processing -> mapping sites.
// The function is concurrency safe. The worker pool size is determined upon program initialization.
func (c *config) Crawl(rawCurUrl string) {

	c.ch <- struct{}{}
	defer func() {
		<-c.ch
		c.wg.Done()
	}()

	if c.pagesLen() >= c.maxPages {
		return
	}

	curUrl, err := url.Parse(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while parsing url: %s", c.root))
		return
	}

	if curUrl.Hostname() != c.root.Hostname() {
		return
	}

	normCurUrl, err := NormalizeURL(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while normalizing url: %s", c.root))
		return
	}

	isFirst := c.addVisited(normCurUrl)
	if !isFirst {
		return
	}

	rawHTML, err := GetHtml(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while fetching html: %v", err))
		return
	}

	//fmt.Println(rawHTML)

	nextUrls, err := GetUrlsFromHTML(rawHTML, c.root)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while fetching urls: %v", err))
		return
	}

	for _, nextUrl := range nextUrls {
		c.wg.Add(1)
		go c.Crawl(nextUrl)
	}

}
