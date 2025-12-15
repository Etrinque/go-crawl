package main

import (
	"net/url"
	"strings"
)

func NormalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	fullURL := parsedURL.Host + parsedURL.Path

	fullURL = strings.ToLower(fullURL)

	fullURL = strings.TrimPrefix(fullURL, "/")

	return fullURL, nil
}
