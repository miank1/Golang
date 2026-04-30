package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
	}
}
