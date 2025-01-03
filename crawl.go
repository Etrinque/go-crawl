package main

import (
	"fmt"
	"net/url"
	"sync"
)

var errLog []error

type config struct {
	pages map[string]int
	root  *url.URL
	mut   *sync.Mutex
	wg    *sync.WaitGroup
	ch    chan struct{}
}

// NewConfig returns a new concurrency config with X number worker-pool size.
// Channel buffered to worker-pool size
func (c *config) NewConfig(root *url.URL, numWorker int, pages map[string]int) *config {

	config := &config{
		pages: pages,
		root:  root,
		wg:    new(sync.WaitGroup),
		mut:   new(sync.Mutex),
		ch:    make(chan struct{}, numWorker),
	}
	config.wg.Add(numWorker)
	return config
}

// Visited checks if the current page has been visited, false (not visited) by default
func (c *config) Visited(rawCurrentUrl string) bool {
	if c.pages != nil && c.pages[rawCurrentUrl] != 0 {
		return true
	}

	return false
}

// Crawl wip: Convert to config Method and refactor for concurrency
func (c *config) Crawl(rawCurUrl string) {

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

	c.mut.Lock()
	_, ok := c.pages[normCurUrl]
	if ok {
		c.pages[normCurUrl]++
		c.mut.Unlock()
		return
	} else {
		c.pages[normCurUrl] = 1
	}
	c.mut.Unlock()

	rawHTML, err := GetHtml(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while fetching html: %v", err))
		return
	}

	fmt.Println(rawHTML)

	c.mut.Lock()
	nextUrls, err := GetUrlsFromHTML(rawHTML, rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while fetching urls: %v", err))
		return
	}
	c.mut.Unlock()

	for _, nextUrl := range nextUrls {
		c.Crawl(nextUrl)
		//time.Sleep(1 * time.Second)
	}

}
