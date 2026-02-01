// Select With Channel Close Detection

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(300 * time.Millisecond)
			ch <- i
		}
		close(ch)
	}()

	for {
		select {
		case num, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received:", num)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout waiting for data")
		}
	}

}
