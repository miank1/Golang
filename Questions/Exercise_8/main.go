// Buffered vs Unbuffered Channels

package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello"
	}()

	fmt.Println(<-ch)

	close(ch)

	// buffered channel

	ch2 := make(chan string, 2)
	ch2 <- "go"
	ch2 <- "lang"

	close(ch2) // closing of the channel is mandatory

	for x := range ch2 {
		fmt.Println(x)
	}

}
