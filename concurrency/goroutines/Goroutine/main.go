package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello")
}

func work(id int) {
	fmt.Println("Worker ", id)
}

func main() {

	go printHello()
	time.Sleep(1 * time.Second)
	fmt.Println("World")

	for i := 0; i < 10; i++ {
		go work(i)
	}

	time.Sleep(time.Second)
}
