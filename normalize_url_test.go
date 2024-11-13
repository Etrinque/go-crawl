package main

import (
	"fmt"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	var tests = []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "strip prefix/https, (/) postfix",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "strip prefix/https (/) postfix",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "strip prefix/http",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "strip prefix/http",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
	}
	// map urls to single normalized outputs blog.boot.dev/path

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NormalizeURL(tt.inputURL)
			if err != nil {
				t.Errorf("Test %v, - '%s' FAILED: unexpected error: %v", i, tt.name, err)
				return
			}
			if actual != tt.expected {
				t.Errorf("Test %v - %s FAILED: expectedResult != %sgo test .", i, tt.name, tt.inputURL)
				return
			}

			t.Logf("test: %v actual: %s == expected: %s", i, actual, tt.expected)

		})
	}
}

func TestGetUrlsFromHTML(t *testing.T) {
	var tests = []struct {
		name      string
		inputBody string
		inputUrl  string
		expected  []string
	}{
		{
			name:      "absolute and relative urls",
			inputBody: "\n<html>\n\t<body>\n\t\t<a href=\"/path/one\">\n\t\t\t<span>Boot.dev</span>\n\t\t</a>\n\t\t<a href=\"https://other.com/path/one\">\n\t\t\t<span>Boot.dev</span>\n\t\t</a>\n\t</body>\n</html>\n",
			inputUrl:  "https://blog.boot.dev",
			expected:  []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:      "absolute and relative urls",
			inputBody: "",
			inputUrl:  "",
			expected:  []string{"", ""},
		},
		{
			name:      "absolute and relative urls",
			inputBody: "",
			inputUrl:  "",
			expected:  []string{"", ""},
		},
	}

	for _, t := range tests {
		lnk, err := GetUrlsFromHTML(t.inputBody)
		if err != nil {
			fmt.Errorf("Test %v, - '%s' FAILED: unexpected error: %v", t.name, t.inputBody, err)
		}
		for i := range lnk {
			if lnk[i] != t.expected[i] {
				fmt.Errorf("Test %v, - '%s' FAILED: expectedResult != %sgo test .", t.name, t.inputBody, t.expected[i])
			}
		}

	}
}
