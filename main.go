package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	var BASE_URL string

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

	raw, err := GetHtml(BASE_URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(raw)
}
