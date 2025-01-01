package main

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

var errLog []error

func Crawl(rawUrl, rawCurUrl string, pages map[string]int) map[string]int {

	curUrl, err := url.Parse(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while normalizing url: %s", rawUrl))
		return pages
	}

	baseUrl, err := url.Parse(rawUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while parsing url: %s", rawUrl))
		return pages
	}

	if curUrl.Hostname() != baseUrl.Hostname() {
		return pages
	}

	normCurUrl, err := NormalizeURL(rawCurUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while normalizing url: %s", rawUrl))
		return pages
	}

	_, ok := pages[normCurUrl]
	if ok {
		pages[normCurUrl]++
		return pages
	}

	pages[normCurUrl] = 1

	rawHTML, err := GetHtml(normCurUrl)
	if err != nil {
		err = errors.New("error while fetching html")
		errLog = append(errLog, err)
	}
	fmt.Println(rawHTML)
	time.Sleep(500 * time.Millisecond)

	nextUrls, err := GetUrlsFromHTML(normCurUrl)
	if err != nil {
		err := errors.New("error while fetching urls")
		errLog = append(errLog, err)
	}

	for _, nextUrl := range nextUrls {
		Crawl(rawUrl, nextUrl, pages)
	}

	return pages
}
