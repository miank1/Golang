package main

import "fmt"

func func1(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch)
}

func func2(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i * 10
	}
	close(ch)
}

func mergeChannels(ch1, ch2 <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		for val := range ch1 {
			out <- val
		}

	}()

	go func() {
		for val := range ch2 {
			out <- val
		}

	}()

	return out
}

func main() {

	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	go func1(ch1)
	go func2(ch2)

	out := mergeChannels(ch1, ch2)

	for i := 1; i <= 10; i++ {
		fmt.Println(<-out)
	}

}
