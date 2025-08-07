package main

import (
	"fmt"
	"time"
)

func firstSlow() int {
	time.Sleep(500 * time.Millisecond)
	return 1
}

func secondSlow() int {
	time.Sleep(100 * time.Millisecond)
	return 2
}

func firstRacer() int {
	ch := make(chan int)

	go func() {
		ch <- firstSlow()
	}()

	go func() {
		ch <- secondSlow()
	}()

	return <-ch
}

func main() {

	result := firstRacer()
	fmt.Println("First Result", result)
}
