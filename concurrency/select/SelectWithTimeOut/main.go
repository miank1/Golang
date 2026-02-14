package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	// ch <- "Hello !!!" // If you don't receive a message on channel for a while then case with timer gets executed.
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: no message received")
	}
	close(ch)
}
