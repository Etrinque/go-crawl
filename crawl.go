package main

import (
	"fmt"
	"net/url"
	"sync"
)

var errLog []error

type concurrent struct {
	pages map[string]int
	root  *url.URL
	mut   *sync.Mutex
	wg    *sync.WaitGroup
	ch    chan struct{}
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

	rawHTML, err := GetHtml(normCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while fetching html: %v", err))
		return
	}

	fmt.Println(rawHTML)

	nextUrls, err := GetUrlsFromHTML(rawHTML, normCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while fetching urls: %v", err))
		return
	}

	for _, nextUrl := range nextUrls {
		Crawl(rawBaseUrl, nextUrl, pages)
	}

}
