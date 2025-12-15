package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// getURLsFromHTML extracts all URLs from anchor tags in the HTML body.
// Relative URLs are resolved against the base URL.
// Returns a slice of absolute URL strings.

func getUrlsFromHTML(htmlBody string, baseUrl *url.URL) ([]string, error) {

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("failure to parse html body. error: %w", err))
		return nil, err
	}

	var urls []string
	var extractUrls func(*html.Node)

	extractUrls = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						// Log error and skip
						errLog = append(errLog, fmt.Errorf("error parsing href: %w", err))
						continue
					}
					resolvedUrls := baseUrl.ResolveReference(href)
					urls = append(urls, resolvedUrls.String())
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			extractUrls(child)
		}
	}
	extractUrls(doc)

	return urls, nil
}
