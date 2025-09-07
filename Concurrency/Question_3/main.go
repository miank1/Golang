// You have two writer goroutines, each writing numbers into the same channel.

// First goroutine writes 1 to 5.

// Second goroutine writes 6 to 10.

// A reader goroutine should read and print numbers until all are done.

// ⚠️ Hint: You’ll need to think about when to close the channel so the reader exits properly.

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
