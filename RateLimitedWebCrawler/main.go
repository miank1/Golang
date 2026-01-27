package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// fetchURL tries to GET a URL and returns it if status == 200
func fetchURL(url string, wg *sync.WaitGroup, sem chan struct{}, results chan<- string) {
	defer wg.Done()

	// limit concurrency
	sem <- struct{}{}
	defer func() { <-sem }() // release slot

	// Make HTTP request with timeout
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("❌ Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Collect only successful responses
	if resp.StatusCode == http.StatusOK {
		results <- url
		fmt.Printf("✅ Success: %s\n", url)
	} else {
		fmt.Printf("⚠️ Skipped (status %d): %s\n", resp.StatusCode, url)
	}
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/delay/2",

		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/delay/2",

		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/delay/2",
		// ... add up to 100 URLs
	}

	const concurrencyLimit = 10

	sem := make(chan struct{}, concurrencyLimit) // semaphore channel
	results := make(chan string, len(urls))      // store successful URLs
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg, sem, results)
	}

	wg.Wait()
	close(results)

	fmt.Println("\n---- Successful URLs ----")
	for url := range results {
		fmt.Println(url)
	}
}
