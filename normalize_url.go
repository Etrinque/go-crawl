package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
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


func getUrlsFromHTML(htmlBody string) ([]string, error) {
	var linkNodes []string

	htmlReader := strings.NewReader(htmlBody)
	node, err := html.Parse(htmlReader)
	if err != nil {
		return []string{}, fmt.Errorf("failure to parse body. error: %v", err)
		
	}

	var recurse func(*html.Node) 
	recurse = func(n *html.Node){
		
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

func clientReq(url string) error {
	var reader strings.Reader
	var client *http.Client

	req, err := http.NewRequest("GET", url, &reader)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Request failure: %v", err)
	}

	

	getUrlsFromHTML()

	return nil
}