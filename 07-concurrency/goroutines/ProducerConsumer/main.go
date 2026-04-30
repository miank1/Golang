package main

import "fmt"

func main() {

	ch := make(chan int)

	// Producer
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}

		close(ch)
	}()

	// consumer
	for val := range ch {
		fmt.Println(val)
	}
}
