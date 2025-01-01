package main

import (
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"
)

var errLog []error

type concurrent struct {
	pages map[string]int
	root  *url.URL
	mut   *sync.Mutex
	wg    *sync.WaitGroup
	ch    chan struct{}
}

func Crawl(rawUrl, rawCurUrl string, pages map[string]int) {

	curUrl, err := url.Parse(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while normalizing url: %s", rawUrl))
		return
	}

	baseUrl, err := url.Parse(rawUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while parsing url: %s", rawUrl))
		return
	}

	if curUrl.Hostname() != baseUrl.Hostname() {
		return
	}

	normCurUrl, err := NormalizeURL(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while normalizing url: %s", rawUrl))
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
		err = errors.New("error while fetching html")
		errLog = append(errLog, err)
	}
	fmt.Println(rawHTML)
	time.Sleep(500 * time.Millisecond)

	nextUrls, err := GetUrlsFromHTML(rawHTML, normCurUrl)
	if err != nil {
		err := errors.New("error while fetching urls")
		errLog = append(errLog, err)
	}

	for _, nextUrl := range nextUrls {
		Crawl(rawUrl, nextUrl, pages)
	}

}
