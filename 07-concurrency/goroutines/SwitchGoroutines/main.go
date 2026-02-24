package main

import (
	"fmt"
	"time"
)

func goroutine1(done chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println("Goroutine 1")
		time.Sleep(500 * time.Millisecond)
		done <- true
	}
}

func goroutine2(done chan bool) {
	for i := 0; i < 5; i++ {
		<-done
		fmt.Println("Goroutine 2")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	done := make(chan bool)
	go goroutine1(done)
	go goroutine2(done)
	done <- true
	time.Sleep(3 * time.Second)
}
