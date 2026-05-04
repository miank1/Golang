package main

import "fmt"

func generator(ch chan int) <-chan int {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

	return ch
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func main() {
	ch := make(chan int, 5)
	gen := generator(ch)
	sq := square(gen)

	for val := range sq {
		fmt.Println(val)
	}
}
