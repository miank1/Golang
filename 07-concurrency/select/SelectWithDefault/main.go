package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	//ch <- "Hello" // If no message is received on channel then 
	select {
	case msg := <-ch:
		fmt.Println("Received :", msg)
	default:
		fmt.Println("No messages received")
	}
}
