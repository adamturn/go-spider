package spider

import (
	"fmt"
	"testing"
)

// TestCrawl calls Crawl with a url string
func TestCrawl(t *testing.T) {
	url := "https://www.federalreserve.gov/releases/h15/"
	fmt.Printf("URL value: %v \n", url)
	fmt.Printf("URL type: %T \n", url)

	body := Crawl(url)
	fmt.Println("Response value:", body)
	fmt.Printf("Response type: %T \n", body)
}

// TestSpinWeb calls SpinWeb with a slice of url strings
func TestSpinWeb(t *testing.T) {
	urls := []string{
		"https://www.federalreserve.gov/releases/h15/",
	}

	data := SpinWeb(urls)
	if data == nil {
		t.Error("Did not receive any data!")
	}
	for i, element := range data {
		fmt.Printf("Element %d: %v", i, element)
	}
}
