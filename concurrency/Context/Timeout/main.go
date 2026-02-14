package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine stopped:", ctx.Err())
		default:
			fmt.Println("Working ...")
		}
	}()

	time.Sleep(3 * time.Second)

	fmt.Println("Main exited")

}
