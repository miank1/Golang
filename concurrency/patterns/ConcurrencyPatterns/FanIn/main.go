// FIXME: func1 and func2 must close channels - missing close() causes goroutine leak / blocks forever on range
package main

import "fmt"

func func1(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}

func func2(ch chan<- int) {
	for i := 5; i <= 8; i++ {
		ch <- i
	}
	close(ch)
}

func fanIn(ch1, ch2 <-chan int) <-chan int {

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

	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	go func1(ch1)
	go func2(ch2)

	out := fanIn(ch1, ch2)

	for i := 0; i < 6; i++ {
		fmt.Println(<-out)
	}

}
