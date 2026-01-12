package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)
	stop := make(chan struct{}) // stop signal

	// Ping Goroutine
	go func() {
		for {
			select {
			case <-stop:
				return
			case msg := <-ping:
				fmt.Println("Ping:", msg)
				pong <- "pong"
			}
		}
	}()

	// Pong Goroutine
	go func() {
		for {
			select {
			case <-stop:
				return
			case msg := <-pong:
				fmt.Println("Pong:", msg)
				ping <- "ping"
			}
		}
	}()

	// Start game
	ping <- "ping"

	// Run for 1 sec
	time.Sleep(1 * time.Second)

	// Send only stop signal
	close(stop)

	// Give goroutines time to exit
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Finished.")
}
