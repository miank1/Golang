package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
	}()

	fmt.Println(<-ch)
}
