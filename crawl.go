package main

import "fmt"

var errors []error

func Crawl(rawUrl, curUrl string, pages map[string]string) []error {

	curUrl, err := NormalizeURL(rawUrl)
	if err != nil {
		errors = append(errors, fmt.Errorf("error while normalizing url"))
	}

	Crawl(rawUrl, curUrl, pages)
}
