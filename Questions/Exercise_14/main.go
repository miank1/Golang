package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		select {
		case <-time.After(3 * time.Second):
			ch <- "data fetched"
		case <-ctx.Done():
			// stop goroutine → no leak
			return
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result := fetchData(ctx)

	select {
	case msg := <-result:
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("Timeout!")
	}
}
