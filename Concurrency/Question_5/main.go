// A goroutine simulates doing some work (e.g., sleep 2 seconds and then send "work done" on a channel).

// The main goroutine should select between:

// Receiving the "work done" message.

// A timeout channel created with time.After(1 * time.Second).

// 👉 Expected behavior:

// If work finishes before timeout → print "Received: work done".

// If timeout happens first → print "Timeout! Work took too long".

package main

import (
	"fmt"
	"time"
)

func DoWork(ch chan string) {

	time.Sleep(2 * time.Second)
	fmt.Println("Working...")
	ch <- "work done"

}

func main() {
	ch := make(chan string)
	go DoWork(ch)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! Work took too long")
	}

}
