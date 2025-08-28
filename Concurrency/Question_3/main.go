package main

import (
	"fmt"
	"sync"
)

func Writer_1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func Writer_2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 6; i <= 10; i++ {
		ch <- i
	}
}

func Reader(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2) // only writers

	// Writers
	go Writer_1(ch, &wg)
	go Writer_2(ch, &wg)

	// Reader (no wg needed, exits when channel closes)
	go Reader(ch)

	// Close channel once all writers finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// wait for writers
	wg.Wait()
}
