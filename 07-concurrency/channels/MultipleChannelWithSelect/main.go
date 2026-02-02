package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- 5
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			ch2 <- "one"
			ch2 <- "two"
			ch2 <- "three"
		}
		close(ch2)
	}()

	for {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				ch1 = nil // Stop receiving from ch1 once it's closed
			} else {
				fmt.Println("Received from ch1:", msg1)
			}

		case msg2, ok := <-ch2:
			if !ok {
				ch2 = nil // Stop receiving from ch2 once it's closed
			} else {
				fmt.Println("Received from ch2:", msg2)
			}
		}

		// Exit
		if ch1 == nil && ch2 == nil {
			break
		}
	}

}
