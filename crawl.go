package main

import (
	"fmt"
	"strings"
)

var errLog []error

func Crawl(rawUrl, curUrl string, pages map[string]int) map[string]int {

	r, err := NormalizeURL(rawUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error while normalizing url: %s", rawUrl))
	}

	// TODO: curUrl to be replaced each depth
	curUrl = r

	// TODO: if curUrl is not of the base domain, log error and move on to next
	if !strings.Contains(curUrl, r) {
		errLog = append(errLog, fmt.Errorf("curUrl out of domain: %s", rawUrl))
		return pages
	}

	_, ok := pages[curUrl]
	if ok {
		pages[curUrl]++
		return pages
	}

	pages[curUrl] = 1
	fmt.Println(GetHtml(curUrl))

	Crawl(r, curUrl, pages)

	return pages
}
