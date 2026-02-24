package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("In Goroutine")
	}()

	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}
