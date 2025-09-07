// Cancel a Goroutine Using Context
// Write a Go program that:

// Starts a worker goroutine.

// Cancels it after 3 seconds using context.WithCancel().

// The worker should print "Working..." every 500ms until it's cancelled, then print "Stopped" and exit.

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped")
			return
		default:
			fmt.Println("Working ...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second)

	cancel()

	time.Sleep(1 * time.Second)
}
