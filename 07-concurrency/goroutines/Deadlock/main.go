package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("Sent 1 to channel")

	go func() {
		fmt.Println(<-ch)
	}()

	time.Sleep(1 * time.Millisecond)
}
