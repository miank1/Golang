// Write a program that concurrently calculates the sum of integers in two slices using goroutines and channels.
// The program should print the sum of each slice and the total sum.

package main

import (
	"fmt"
	"sync"
)

func SumOfIntegers(slice []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	sum := 0

	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}

	ch <- sum
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}

	var wg sync.WaitGroup

	// Create buffered channel
	ch := make(chan int, 2)

	wg.Add(2)

	go SumOfIntegers(slice1, &wg, ch)
	go SumOfIntegers(slice2, &wg, ch)

	wg.Wait()
	close(ch)

	totalSum := 0

	for sum := range ch {
		fmt.Printf("Sum: %d\n", sum)
		totalSum += sum
	}

	fmt.Printf("Total Sum: %d\n", totalSum)
}
