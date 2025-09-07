// Cancel a function Early with context

package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func printNumbers(ctx context.Context) {

	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("stopped:", ctx.Err())
			return
		case <-time.After(1 * time.Second):
			fmt.Println(i)
		}
	}

	fmt.Println("Finished printing all numbers")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	printNumbers(ctx)

}
