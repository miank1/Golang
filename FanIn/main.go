package main

import (
	"fmt"
	"sync"
)

func generate(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan string)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	ch1 := generator(1, 5)
	ch2 := generator(6, 10)
	ch3 := generator(11, 15)

	combined := fanIn(ch1, ch2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-combined)
	}

}
