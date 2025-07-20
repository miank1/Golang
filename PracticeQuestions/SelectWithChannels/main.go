// Explain the select statement and how it is used to manage multiple channels concurrently.
// Provide an example of a program that uses a select statement to handle three channels: one for receiving messages,
// one for sending a periodic signal, and one for handling a timeout using time.After. Demonstrate how to avoid deadlocks
// and handle the case where no channels are ready. Discuss the implications of using a default case in a select statement
// and how it affects performance or behavior.
package main

import (
	"fmt"
	"time"
)

func HandleChannels(ch chan int) {
	ticker := time.NewTicker(500 * time.Millisecond) // Send signal every 500ms
	defer ticker.Stop()

	count := 0
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Received message %d\n", val)
			return // Exit after receiving a message
		case <-time.After(2 * time.Second): // Longer timeout
			fmt.Println("Operation timed out")
			return
		case <-ticker.C:
			count++
			fmt.Printf("Tick %d at %v\n", count, time.Now().Format("15:04:05.000"))
			if count >= 3 { // Stop after 3 ticks
				fmt.Println("Stopping after 3 ticks")
				return
			}
		}
	}
}

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second) // Simulate delayed send
		ch <- 42
	}()
	HandleChannels(ch)
}
