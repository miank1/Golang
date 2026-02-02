package main

import (
	"fmt"
	"time"
)

// squares numbers
func worker(id int, numbers <-chan int, results chan<- int) {
	for n := range numbers {
		fmt.Printf("Worker %d processing %d\n", id, n)
		time.Sleep(500 * time.Millisecond) // simulate some work
		results <- n * n                   // square no
	}
}

func main() {
	numbers := make(chan int, 10)
	results := make(chan int, 10)

	// Start 3 workers (fan-out)
	for w := 1; w <= 3; w++ {
		go worker(w, numbers, results)
	}

	for i := 1; i <= 10; i++ {
		numbers <- i
	}
	close(numbers) // no more numbers

	// Collect results
	for i := 1; i <= 10; i++ {
		fmt.Println("Result:", <-results)
	}
}
