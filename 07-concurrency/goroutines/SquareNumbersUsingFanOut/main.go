// Write a Go program that implements the fanout pattern to process a slice of integers.
// The program should take a slice of integers, distribute the elements across multiple goroutines
// to calculate their squares, and collect the results into a single slice. Ensure proper synchronization
// and handle potential errors.
package main

import (
	"fmt"
	"sync"
)

// SquareResult holds the result of squaring a number or an error
type SquareResult struct {
	Index int
	Value int
	Err   error
}

// square calculates the square of a number and sends the result to a channel
func square(index int, num int, results chan SquareResult) {
	// Simulate potential error condition (e.g., negative numbers not allowed)
	if num < 0 {
		results <- SquareResult{Index: index, Err: fmt.Errorf("negative number %d at index %d", num, index)}
		return
	}
	results <- SquareResult{Index: index, Value: num * num}
}

func main() {
	// Input slice
	numbers := []int{1, 2, 3, 4, 5, 1, 6, 7, 8, 9}

	// Channel to collect results
	results := make(chan SquareResult, len(numbers))

	// WaitGroup for synchronization
	var wg sync.WaitGroup

	// Fanout: Distribute work to goroutines
	for i, num := range numbers {
		wg.Add(1)
		go func(index, value int) {
			defer wg.Done()
			square(index, value, results)
		}(i, num)
	}

	// Close results channel after all goroutines complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	output := make([]int, len(numbers))
	hasError := false
	for result := range results {
		if result.Err != nil {
			fmt.Println("Error:", result.Err)
			hasError = true
			continue
		}
		output[result.Index] = result.Value
	}

	// Print results if no errors
	if !hasError {
		fmt.Println("Squared numbers:", output)
	} else {
		fmt.Println("Processing completed with errors.")
	}
}
