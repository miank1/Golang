// Rate Limiting with Ticker + Select

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for range ticker.C {
			req, ok := <-requests
			if !ok {
				fmt.Println("All requests processed")

				return
			}
			fmt.Println("Processing request:", req)
		}
	}()

	wg.Wait()

}
