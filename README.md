# Golang HTTP Spider
For basic web crawling. Work in progress.

## Quickstart
```go
package main

import (
    "fmt"

    spider "github.com/adamturn/go-spider"
)

func main() {
    // define a url string
    url := "https://example.site.com"
    // Crawl handles standard errors and
    // returns response body as a string
    page := spider.Crawl(url)
    // print HTML to stdout and
    // parse with other service
    fmt.Println(url)
}
```

## Test
$ go test -v
