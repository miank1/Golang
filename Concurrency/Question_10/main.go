// Generate squares and print them - pipeline
package main

import (
	"fmt"
)

func Generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- i
		}
	}()
	return out
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for x := range in {
			out <- x * x
		}
	}()
	return out
}

// Stage 3: Filter even squares
func FilterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for x := range in {
			if x%2 == 0 {
				out <- x
			}
		}
	}()
	return out
}

func Print(ch <-chan int) {
	for msg := range ch {
		fmt.Printf("%d ", msg)
	}
}

func main() {
	nums := Generate(10)
	squares := Square(nums)
	evens := FilterEven(squares)
	Print(evens)
}
