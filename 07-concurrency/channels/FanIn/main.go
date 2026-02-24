package main

import (
	"fmt"
	"time"
)

// Producer function
func producer(name string, out chan<- string) {
	for i := 1; i <= 5; i++ {
		out <- fmt.Sprintf("%s: message %d", name, i)
		time.Sleep(time.Millisecond * 500) // Simulate work
	}
	close(out)
}

// Fan-In function: Merges multiple channels into one
func fanIn(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)

	// Use WaitGroup to ensure both goroutines finish before closing out
	go func() {
		for val := range ch1 {
			out <- val
		}
	}()

	go func() {
		for val := range ch2 {
			out <- val
		}
	}()

	// Close output channel when both input channels are finished
	go func() {
		time.Sleep(time.Second * 3)
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two producers
	go producer("Producer 1", ch1)
	go producer("Producer 2", ch2)

	// Merge both channels
	merged := fanIn(ch1, ch2)

	// Read from the merged channel
	for msg := range merged {
		fmt.Println(msg)
	}
}
