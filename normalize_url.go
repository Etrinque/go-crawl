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

//func clientReq(url string) error {
//	var reader strings.Reader
//	var client *http.Client
//
//	req, err := http.NewRequest("GET", url, &reader)
//	if err != nil {
//		return fmt.Errorf("error making request: %v", err)
//	}
//
//	resp, err := client.Do(req)
//	if err != nil {
//		return fmt.Errorf("Request failure: %v", err)
//	}
//
//	if resp.StatusCode >= 299 || resp.StatusCode <= 199 {
//		return fmt.Errorf("Request failed with status code %d", resp.StatusCode)
//	}
//
//	return nil
//}
