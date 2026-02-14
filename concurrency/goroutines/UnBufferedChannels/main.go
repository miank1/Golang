package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	fmt.Println("Sending a message")
	ch <- 42 // Send a value on the channel
	fmt.Println("Message sent")
}

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	go sender(ch) // Start the sender goroutine

	fmt.Println("Waiting to receive message")
	time.Sleep(2 * time.Second) // Simulate some work

	// Receive the value from the channel
	fmt.Println("Receiving a message")
	value := <-ch
	fmt.Println("Received message:", value)
}
