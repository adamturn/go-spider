package spider

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Crawl requests url and returns HTTP response body
func Crawl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

// SpinWeb iterates over passed urls and Crawls each of them
func SpinWeb(urls []string) []string {
	var pages []string
	for _, url := range urls {
		pages = append(pages, Crawl(url))
	}

	return pages
}
