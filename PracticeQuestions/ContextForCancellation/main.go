// explain the concept of context (context.Context) and how it is used for cancellation and timeouts in concurrent programs.
// Provide an example demonstrating how to use context.WithTimeout to cancel a long-running goroutine operation after a
// specific duration, including proper error handling and cleanup.

package main

import (
	"context"
	"fmt"
	"time"
)

func LongRunningOperation(ctx context.Context, taskNo int) error {

	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("task %d completed successfully\n", taskNo)
		return nil

	case <-ctx.Done():
		fmt.Printf("Task %d cancelled: %v\n", taskNo, ctx.Err())
		return ctx.Err()
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // changine the time stamp here will work.

	defer cancel()

	go func() {
		if err := LongRunningOperation(ctx, 1); err != nil {
			fmt.Printf("Goroutine error: %v\n", err)
			return
		}
		fmt.Println("Goroutine Finished")
	}()

	<-ctx.Done()
	fmt.Println("Main: Context done with error:", ctx.Err())

	// Give a moment to see goroutine output
	time.Sleep(100 * time.Millisecond)

}
