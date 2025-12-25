package util

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetHtml(rawUrl string) (string, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(rawUrl)
	if err != nil {
		return "", fmt.Errorf("HTTP request failure: %w", err)

	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// verify content type
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(strings.ToLower(contentType), "text/html") {
		return "", fmt.Errorf("expected -> text/html | got -> %s", contentType)
	}

	//Return the webpage's HTML if successful
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
