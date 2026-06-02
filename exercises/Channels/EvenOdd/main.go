package main

import (
	"fmt"
	"sync"
)

func Even(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= 10; i += 2 {
		<-ch
		fmt.Println(i)
		ch <- 1
	}

}

func Odd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i += 2 {
		ch <- 1
		fmt.Println(i)
		<-ch

	}
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go Even(ch, &wg)
	go Odd(ch, &wg)

	wg.Wait()
	close(ch)

}
