package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// GetUrlsFromHTML reads in the current raw HTML body.
// Recursively dives through pages for <a href> tags -> child nodes.

func GetUrlsFromHTML(htmlBody string, baseUrl *url.URL) ([]string, error) {
	var linkNodes []string

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("failure to parse body. error: %w", err))
		return nil, err
	}

	var recurse func(*html.Node)
	recurse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						errLog = append(errLog, fmt.Errorf("error parsing href: %w", err))
						continue
					}
					resolvd := baseUrl.ResolveReference(href)
					linkNodes = append(linkNodes, resolvd.String())
				}
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			recurse(child)
		}
	}
	recurse(doc)

	//for _, elem := range linkNodes {
	//	fmt.Println(elem)
	//}

	return linkNodes, nil
}
