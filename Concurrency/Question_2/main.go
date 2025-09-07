// Write a Go program where:

// One goroutine generates numbers from 1 to 10 and sends them into a channel.

// Another goroutine reads from the channel and prints the numbers.

// Make sure the program exits cleanly after all numbers are printed (no deadlocks, no infinite waiting).

package main

import "fmt"

func Write(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}

	close(ch)
}

func Read(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}
func main() {
	ch := make(chan int)
	go Write(ch)
	Read(ch)
}
