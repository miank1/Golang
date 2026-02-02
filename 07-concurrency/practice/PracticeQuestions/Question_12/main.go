// Basic Goroutine Synchronization

package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d done \n ", id)
}
func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
}
