package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Set a deadline exactly 2 seconds from now
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine stopped:", ctx.Err())
		}
	}()

	// Simulate some work
	time.Sleep(3 * time.Second)

	fmt.Println("Main exited")
}
