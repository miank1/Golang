package main

import "fmt"

func main() {
	ch := make(chan int)
	// this will cause a deadlock since its a unbuffered channel
	// and send operations is blocks until receiver is ready

	go func() {
		fmt.Println(<-ch)
	}()

	ch <- 1
}
