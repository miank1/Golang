package main

import (
	"fmt"
	"time"
)

// Produce function simulates a producer generating items and sending them to the channel
func produce(ch chan<- int, id int) {
	for i := 0; ; i++ {
		fmt.Printf("Producer %d: producing item %d\n", id, i)
		ch <- i
		time.Sleep(time.Millisecond * 500) // simulate time taken to produce an item
	}
}

// Consume function simulates a consumer receiving items from the channel
func consume(ch <-chan int, id int) {
	for {
		item := <-ch
		fmt.Printf("Consumer %d: consuming item %d\n", id, item)
		time.Sleep(time.Second) // simulate time taken to consume an item
	}
}

func main() {
	buffer := make(chan int, 10) // buffered channel with a capacity of 10

	// Start producer goroutines
	for i := 0; i < 3; i++ {
		go produce(buffer, i)
	}

	// Start consumer goroutines
	for i := 0; i < 2; i++ {
		go consume(buffer, i)
	}

	// Let the goroutines run for a while
	time.Sleep(time.Second * 3)
}
