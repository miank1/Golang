// Write a Go program that: Generates numbers 1 to 20. Uses 3 concurrent workers to square the numbers.
// Merges their results. Prints them.

package main

import (
	"fmt"
	"sync"
)

// Generate numbers 1..n
func Generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- i
		}
	}()
	return out
}

// Worker squares numbers
func Worker(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		fmt.Printf("Worker %d processing %d\n", id, num)
		out <- num * num
	}
}

func main() {
	// Step 1: Generate numbers
	numbers := Generate(20)

	results := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Worker(i, numbers, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Result:", res)
	}
}
