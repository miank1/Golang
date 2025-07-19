// Write a Go program that:

// Has two goroutines:

// One sends "ping" to a channel every 500ms.

// Another sends "pong" to a second channel every 700ms.

// Use a select statement in the main goroutine to:

// Print whichever message arrives first.

// Run this for about 3 seconds using time.After.

package main

import "time"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// first goroutine
	go func() {
		for {
			ch1 <- "ping"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// second goroutine
	go func() {
		for {
			ch2 <- "pong"
			time.Sleep(700 * time.Millisecond)
		}
	}()

	timeout := time.After(10 * time.Second)

	for {
		select {
		case msg := <-ch1:
			println("Received:", msg)
		case msg := <-ch2:
			println("Received:", msg)
		case <-timeout:
			println("Timeout after 3 seconds")
			return
		}
	}
}
