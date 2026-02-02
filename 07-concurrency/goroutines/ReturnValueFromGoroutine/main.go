package main

import "fmt"

func main() {
	result := make(chan int, 1)
	go sum(2, 3, result)
	value := <-result

	fmt.Printf("Value: %d\n", value)
	close(result)
}

func sum(a, b int, result chan int) {
	sum := a + b
	result <- sum
}
