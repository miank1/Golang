package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "done"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")

	}

}
