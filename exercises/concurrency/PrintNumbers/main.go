package main

import (
	"fmt"
	"sync"
)

func PrintNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func main() {

	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	var wg sync.WaitGroup

	wg.Add(1)

	go PrintNumbers(ch, &wg)

	wg.Wait()
}
