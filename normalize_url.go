package main

import (
	"strings"
)

func NormalizeURL(url string) (string, error) {
	var normURL string
	var normalized string

	if strings.HasPrefix(url, "https://") {
		normURL = strings.TrimPrefix(url, "https://")
	} else if strings.HasPrefix(url, "http://") {
		normURL = strings.TrimPrefix(url, "http://")
	}

	if strings.HasSuffix(normURL, "/") {
		normalized = strings.TrimSuffix(normURL, "/")
		return normalized, nil
	}
	return normURL, nil
}