package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func GetUrlsFromHTML(htmlBody, baeUrl string) ([]string, error) {
	var linkNodes []string

	baseUrl, err := url.Parse(baeUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error parsing base url %v", baseUrl))
		return nil, err
	}

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("failure to parse body. error: %v", err))
		return nil, err
	}

	var recurse func(*html.Node)
	recurse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						errLog = append(errLog, fmt.Errorf("error parsing href: %v", err))
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
