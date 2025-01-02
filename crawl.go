package main

import (
	"fmt"
	"net/url"
	"sync"
	"time"
)

var errLog []error

type config struct {
	pages map[string]int
	root  *url.URL
	mut   *sync.Mutex
	wg    *sync.WaitGroup
	ch    chan struct{}
}

// init new config with X number worker-pool size. Chan buffered to worker-pool size
func newConfig(root *url.URL, numWorker int, pages map[string]int) *config {

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

func Crawl(rawBaseUrl, rawCurUrl string, pages map[string]int) {

	curUrl, err := url.Parse(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while parsing url: %s", rawBaseUrl))
		return
	}

	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while parsing url: %s", rawBaseUrl))
		return
	}

	if curUrl.Hostname() != baseUrl.Hostname() {
		return
	}

	normCurUrl, err := NormalizeURL(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while normalizing url: %s", rawBaseUrl))
		return
	}

	_, ok := pages[normCurUrl]
	if ok {
		pages[normCurUrl]++
		return
	}

	pages[normCurUrl] = 1

	rawHTML, err := GetHtml(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while fetching html: %v", err))
		return
	}

	fmt.Println(rawHTML)

	nextUrls, err := GetUrlsFromHTML(rawHTML, rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while fetching urls: %v", err))
		return
	}

	for _, nextUrl := range nextUrls {
		Crawl(rawBaseUrl, nextUrl, pages)
		time.Sleep(500 * time.Millisecond)

	}

}
