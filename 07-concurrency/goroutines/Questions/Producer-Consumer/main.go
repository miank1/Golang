package main

import (
	"fmt"
	"sync"
)

func Producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

func Consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range ch {
		fmt.Println(x)
	}
}

func main() {

	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go Producer(ch, &wg)
	go Consumer(ch, &wg)

	wg.Wait()

}
