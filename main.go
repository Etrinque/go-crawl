package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	var baseUrl string
	var pages = make(map[string]int)
	var c *config

	if len(args) < 4 {
		fmt.Println("no website provided")
		return
		//os.Exit(1)
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		return
		//os.Exit(1)
	} else {
		baseUrl = args[1]
		fmt.Printf("Starting Crawl for: %s...\n ", baseUrl)
	}

	numWorkers, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Invalid number of workers provided: %v\n", err)
	}

	maxPages, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Printf("Invalid number of pages provided: %v\n", err)
	}

	root, err := url.Parse(baseUrl)
	if err != nil {
		errLog = append(errLog, fmt.Errorf("error parsing root: %v", err))
	}

	config := c.NewConfig(root, numWorkers, maxPages, pages)

	// FIXME: Deadlocking!
	go func() {
		defer config.wg.Done()
		config.Crawl(root.String())
	}()
	config.wg.Wait()

	if errLog != nil {
		for i, err := range errLog {
			i++
			fmt.Printf("Error #%d: %v\n", i, err)
		}
	}

	fmt.Printf("done crawling from root: %s ", c.root.String())

	for k, v := range pages {
		fmt.Printf("Results Page: %s, Occurences: %d\n", k, v)
	}
}
