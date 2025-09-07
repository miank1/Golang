package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)

	fmt.Println("Main: canceling context")
	cancel()

	time.Sleep(1 * time.Second)

}

func worker(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: received cancel signal, exiting.")
			return
		default:
			fmt.Println("Working: doing work")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
