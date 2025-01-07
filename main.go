package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	var c *config

	if len(args) < 4 {
		fmt.Println("no website provided")
		fmt.Println("usage: ./crawler <numWorkers> <maxPages>")
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
	}

	numWorkers, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Invalid number of workers provided: %v\n", err)
	}

	maxPages, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Printf("Invalid number of pages provided: %v\n", err)
	}

	root, err := url.Parse(args[1])
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error parsing root: %v", err))
	}
	fmt.Printf("Starting Crawl for: %s...\n ", root)

	config := c.NewConfig(root, numWorkers, maxPages)

	config.wg.Add(1)
	go config.Crawl(root.String())
	config.wg.Wait()

	fmt.Println("done crawling from root:\t", root)
	fmt.Printf("=============================\n  REPORT for %s\n=============================\n", root.String())

	var tempMap []Page

	for k, v := range config.pages {
		page := Page{k, v}
		tempMap = append(tempMap, page)
	}

	sortedMap := MergeSort(tempMap)
	for _, page := range sortedMap {
		fmt.Printf("Found:\t%d internal links to %s\n", page.val, page.url)
	}
	fmt.Printf("crawled numPages:\t%d\n", config.pagesLen())

	//if errLog != nil {
	//	for i, err := range errLog {
	//		i++
	//		fmt.Printf("Error found #%d:\t%v\n", i, err)
	//	}
	//}
}
