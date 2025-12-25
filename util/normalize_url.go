package util

import (
	"net/url"
	"strings"
)

func NormalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	normalized := parsedURL.Host + parsedURL.Path

	normalized = strings.ToLower(normalized)

	normalized = strings.TrimPrefix(normalized, "/")

	return normalized, nil
}
