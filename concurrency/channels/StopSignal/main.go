package main

import (
	"fmt"
	"time"
)

// How to send a stop channel to a golang channel

func main() {
	fmt.Println("From Main function")

	ch := make(chan bool)
	go func() {
		for {
			select {
			case <-ch:
				return
			default:
				fmt.Println("This is a test routine")
				time.Sleep(time.Second * 3)
			}

		}
	}()

	time.Sleep(time.Second * 10)

	// Stop a go routine
	ch <- true
	fmt.Println("Go routine has stopped")

	// test if the gorountine stopped or not
	time.Sleep(time.Second * 5)
	fmt.Println("This is the end of main func, goroutine no longer run!")

}
