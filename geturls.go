package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func GetUrlsFromHTML(htmlBody string) ([]string, error) {
	var linkNodes []string

	htmlReader := strings.NewReader(htmlBody)
	node, err := html.Parse(htmlReader)
	if err != nil {
		err = fmt.Errorf("failure to parse body. error: %v", err)
		errLog = append(errLog, err)
		return []string{}, err
	}

	var recurse func(*html.Node)
	recurse = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			linkNodes = append(linkNodes, n.Data)
		}
		if n.NextSibling == nil {
			return
		}
		recurse(n.NextSibling)
	}
	recurse(node)

	for _, elem := range linkNodes {
		fmt.Println(elem)
	}

	return linkNodes, nil
}
