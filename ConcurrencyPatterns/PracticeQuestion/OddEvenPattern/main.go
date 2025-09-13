package main

import "fmt"

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func even(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func square(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	nums := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	evenNumbers := even(nums)
	squares := square(evenNumbers)

	for r := range squares {
		fmt.Println(r)
	}
}
