package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetHtml(rawUrl string) (string, error) {
	var rawHTML string

	//Use http.Get to fetch the webpage of the rawURL
	//Return an error if the HTTP status code is an error-level code (400+)
	resp, err := http.Get(rawUrl)
	if err != nil || resp.StatusCode > 399 {
		errLog = append(errLog, fmt.Errorf("get raw url failed: %v", err))
		if resp.StatusCode > 399 {
			errLog = append(errLog, fmt.Errorf("status Code: %d", resp.StatusCode))
		}
		return "", err
	}
	defer resp.Body.Close()

	//Return an error if the response content-type header is not text/html
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		errLog = append(errLog, errors.New("content type is not text/html"))
		return "", err
	}

	//Return any other possible errors
	//Return the webpage's HTML if successful
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		errLog = append(errLog, errors.New("error reading response body"), err)
		return "", err
	}
	rawHTML = string(b)

	return rawHTML, nil

}
