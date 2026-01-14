// Select + Timeout + Loop

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "done"
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println("Received: ", msg)
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Waiting ...")
		}

	}
}
