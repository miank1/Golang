package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 1)
	ch <- 24

	go func() {
		fmt.Println("Received:", <-ch)
	}()

	time.Sleep(1 * time.Second)

}
