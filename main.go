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
	for k, v := range config.pages {
		fmt.Printf("Results Page:\t%s, Occurences:\t%d\n", k, v)
	}
	fmt.Printf("crawled numPages:\t%d", config.pagesLen())

	if errLog != nil {
		for i, err := range errLog {
			i++
			fmt.Printf("Error found #%d:\t%v\n", i, err)
		}
	}
}
