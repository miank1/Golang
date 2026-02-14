package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)

	for event := range ch {
		fmt.Println(event)
	}
}

func sendData(ch chan string) {
	ch <- "Hello"
	ch <- "World"
	ch <- "!"
	close(ch)
}
