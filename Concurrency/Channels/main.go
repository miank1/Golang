package main

import "fmt"

func worker(ch chan int) {
	ch <- 42
}

func main() {

	ch := make(chan int)

	go worker(ch)

	value := <-ch

	fmt.Println(value)
}
