// Implement a Timeout for a Long-Running Task

package main

import (
	"fmt"
	"sync"
	"time"
)

func longTask() {
	time.Sleep(3 * time.Second)
	fmt.Println("Task finished")
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	done := make(chan struct{})

	go func() {
		defer wg.Done()
		longTask()
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Completed within time")
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out!")
	}

}
