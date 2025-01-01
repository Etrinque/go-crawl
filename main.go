package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	var baseUrl string
	var pages = make(map[string]int)

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		baseUrl = args[1]
		fmt.Printf("starting crawl of: %s...\n ", baseUrl)
	}

	//root, err := url.Parse(baseUrl)
	//if err != nil {
	//	errLog = append(errLog, fmt.Errorf("error parsing root: %v", err))
	//}

	//concurrent := &concurrent{root: root, pages: pages, ch: make(chan struct{})}
	//
	//go func() {
	//
	//}()

	pages = Crawl(baseUrl, baseUrl, pages)

	if errLog != nil {
		for i, err := range errLog {
			i++
			fmt.Printf("Error #%d: %v\n", i, err)
		}
	}

	fmt.Println("done crawling")

	for k, v := range pages {
		fmt.Printf("Results Page: %s, Occurences: %d\n", k, v)
	}
}
