// Even-Odd Number Producer-Consumer
// Problem Statement:

// You need to implement two producers:

// One producer generates even numbers from 0 up to N.

// The other producer generates odd numbers from 1 up to N.

// There is one consumer that reads numbers from the channel and prints them.
package main

import (
	"fmt"
	"sync"
)

func Prodcuer_1(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			ch <- i
			fmt.Println("Producer Even Number ", i)
		}
	}
}

func Prodcuer_2(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			ch <- i
			fmt.Println("Producer Odd Number ", i)
		}
	}
}

func Consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range ch {
		fmt.Println("Consumer consumed", x)
	}
}

func EvenOddProducerConsumer(N int) {
	ch := make(chan int, 10)

	var wgProducers sync.WaitGroup
	var wgConsumer sync.WaitGroup

	wgProducers.Add(2)
	wgConsumer.Add(1)

	go Prodcuer_1(N, ch, &wgProducers)
	go Prodcuer_2(N, ch, &wgProducers)
	go Consumer(ch, &wgConsumer)

	wgProducers.Wait()
	close(ch)
	wgConsumer.Wait()
}

func main() {
	EvenOddProducerConsumer(10)
}
