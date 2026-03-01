package main

import (
	"errors"
	"fmt"
	"time"
)

func slowTask() int {
	time.Sleep(2 * time.Second)
	return 42
}

func RunWithTimeout(timeout time.Duration) (int, error) {
	ch := make(chan int, 1)

	go func() {
		result := slowTask()
		ch <- result
	}()

	select {
	case res := <-ch:
		return res, nil

	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func main() {

	// Case 1: Timeout shorter than task duration
	result, err := RunWithTimeout(1 * time.Second)
	fmt.Println("Result:", result, "Error:", err)

	// Case 2: Timeout longer than task duration
	result, err = RunWithTimeout(3 * time.Second)
	fmt.Println("Result:", result, "Error:", err)
}
