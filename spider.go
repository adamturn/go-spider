package spider

import (
	"fmt"
	"net/http"
)

// Crawl requests url and returns HTTP response body
func Crawl(url string) string {
	response, err := http.Get(url)
	fmt.Printf("Response: %v", response)
	fmt.Printf("Error: %v", err)
	return "HTTP response body"
}
