package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/etrinque/go-crawl/util"
)

func main() {

	args := os.Args

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
		// TODO: Refactor/Replace errLog with Logger api
		//errLog = append(errLog, fmt.Errorf("error parsing root: %v", err))
		panic(err)
	}

	fmt.Printf("Starting Crawl for: %s...\n ", root)

	config := newConfig(root, numWorkers, maxPages)

	config.wg.Add(1)
	go config.Crawl(root.String())
	config.wg.Wait()

	fmt.Println("done crawling from root:\t", root)
	var hashes = ""
	var msg = "Root for: "
	for range len(root.String()) + len(msg) {
		hashes += "#"
	}
	fmt.Printf("%s\n"+
		"%s%s"+
		"\n%s\n", hashes, msg, root.String(), hashes)

	var tempMap []util.Page

	for k, v := range config.pages {
		page := util.Page{Url: k, Count: v}
		tempMap = append(tempMap, page)
	}

	sortedMap := util.MergeSort(tempMap)
	for _, page := range sortedMap {
		fmt.Printf("Found:\t%d internal links to %s\n", page.Count, page.Url)
	}
	fmt.Printf("crawled numPages:\t%d\n", config.pagesLen())

	//if errLog != nil {
	//	for i, err := range errLog {
	//		i++
	//		fmt.Printf("Error found #%d:\t%v\n", i, err)
	//	}
	//}
}
