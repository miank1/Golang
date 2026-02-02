package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Go routine cancelled:", ctx.Err())
				return
			default:
				fmt.Println("Working ...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main exited")
}
