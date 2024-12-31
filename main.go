package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	var BASE_URL string
	var pages = make(map[string]int)

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		BASE_URL = args[1]
		fmt.Printf("starting crawl of: %s...\n ", BASE_URL)
	}

	pages = Crawl(BASE_URL, BASE_URL, pages)

	for page, _ := range pages {
		_, err := GetHtml(page)
		if err != nil {
			fmt.Println(err)
		}
	}
}
