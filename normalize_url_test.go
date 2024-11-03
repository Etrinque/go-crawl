package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	var tests = []struct{
		name string
		inputURL string
		expected string
	}{
		{
			name: "strip prefix/https, (/) postfix",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name: "strip prefix/https (/) postfix",
			inputURL: "https://blog.boot.dev/path/", 
			expected: "blog.boot.dev/path",
		},
		{
			name: "strip prefix/http",
			inputURL: "http://blog.boot.dev/path", 
			expected: "blog.boot.dev/path",
		},
		{
			name: "strip prefix/http",
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