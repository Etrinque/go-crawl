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
	node, err := html.Parse(htmlReader)
	if err != nil {
		err = fmt.Errorf("failure to parse body. error: %v", err)
		errLog = append(errLog, err)
		return []string{}, err
	}

	var urls []string
	var recurse func(*html.Node)
	recurse = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						errLog = append(errLog, fmt.Errorf("error parsing href", err))
						continue
					}
					resolvd := baseUrl.ResolveReference(href)
					urls = append(urls, resolvd.String())
				}
			}
			linkNodes = append(linkNodes, n.Data)
		}
		for n = n.FirstChild; n != nil; n = n.NextSibling {
			recurse(n.NextSibling)
		}
	}
	recurse(node)

	//for _, elem := range linkNodes {
	//	fmt.Println(elem)
	//}

	return linkNodes, nil
}
