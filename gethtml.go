package main

import (
	"errors"
	"io"
	"net/http"
)

func GetHtml(rawUrl string) (string, error) {
	var rawHTML string

	//Use http.Get to fetch the webpage of the rawURL
	//Return an error if the HTTP status code is an error-level code (400+)
	resp, err := http.Get(rawUrl)
	if err != nil || resp.StatusCode >= 400 {
		return "", err
	}

	//Return an error if the response content-type header is not text/html
	if resp.Header.Get("Content-Type") != "text/html" {
		return "", errors.New("content type is not text/html")
	}

	//Return any other possible errors
	//Return the webpage's HTML if successful
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	rawHTML = string(b)

	return rawHTML, nil

}
