# Go-Crawl
## A Simple Concurrency-Safe Web-Crawler Implementation in GO!

---
### Table of Contents

* [Overview](#overview)
* [Features](#features)
* [Requirements](#requirements)
* [Installation](#installation)
* [Usage](#usage)

---
## Overview
Go-Crawl is a simple, concurrency-safe web-crawler written in GO. It allows users to crawl websites concurrently and extract internal links.
### Features

Go-Crawl includes the following features:

* **Concurrency Support**: Go-Crawl uses goroutines to crawl multiple URLs concurrently, making it ideal for large-scale web scraping tasks.
* **Robust Error Handling**: The project includes robust error handling mechanisms, including logging and exception handling.
* **Customizable Configuration Options**: Users can customize the crawler's behavior by modifying configuration options.

---
## Requirements

To use Go-Crawl, you need:

* GO 1.22 or later
* `net/url` package
* `golang.org/x/net/html` package

---
## Installation

To install Go-Crawl, run the following commands from your terminal:

* create directory and cd into
```shell
> mkdir project && cd project
```

* clone the repo down
```shell
> git clone https://github.com/etrinque/go-crawl
```

* build the project
```shell
> go buil -O crawl
```

## Usage
From console, run the resulting file with args
```shell

// enter a target URL to start from and number of go routine workers (optional)

> crawl {url} {worker-pool-size}
```
