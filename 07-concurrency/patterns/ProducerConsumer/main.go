package main

import (
	"fmt"
	"sync"
	"time"
)

func Producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 2)
	}

	close(ch)
}

func Consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range ch {
		fmt.Println("Consumed ", x)
	}

}

func main() {

	ch := make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(2)

	go Producer(ch, &wg)
	go Consumer(ch, &wg)

	wg.Wait()
	//close(ch)
}
