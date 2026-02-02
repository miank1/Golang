package main

import (
	"fmt"
	"sync"
)

func Producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 100
	close(ch)
}

func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-ch)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)

	go Producer(ch, &wg)
	go Consumer(ch, &wg)

	wg.Wait()

}
