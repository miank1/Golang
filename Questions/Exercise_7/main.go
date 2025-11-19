// Go Routines & Channels

package main

import "fmt"

func printNumbers(n int, ch chan int) {

	for i := 1; i <= n; i++ {
		ch <- i
	}

	close(ch)
}

func main() {

	ch := make(chan int, 5)
	go printNumbers(5, ch)

	for x := range ch {
		fmt.Println(x)
	}
}
