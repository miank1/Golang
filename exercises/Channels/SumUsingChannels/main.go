// Write a program that: Creates a channel Starts one goroutine Goroutine sends numbers 1 to 5 into the channel
// Main goroutine receives numbers and calculates sum Prints the final sum Must close channel properly Must not deadlock

package main

import "fmt"

func SumNumber(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch)
}

func main() {

	ch := make(chan int, 5)

	go SumNumber(ch)
	total := 0

	for x := range ch {
		total += x

	}

	fmt.Println(total)
}
