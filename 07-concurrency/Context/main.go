package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	fmt.Println("Worker : started")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: cancelled ->", ctx.Err())
			return
		default:
			fmt.Println("Worker: working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("Main: sending cancel signal")

	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main: done")
}
