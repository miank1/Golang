// Write a Go program that creates 3 goroutines. Each goroutine should print numbers from 1 to 5. However, you must ensure that:

// The first goroutine prints all its numbers, ***

// Then the second goroutine prints all its numbers,

// And finally the third goroutine does so.

// Use channels or WaitGroups to coordinate the execution in the correct order.
package main

import (
	"fmt"
	"sync"
)

func First(text string, ch1, ch2 chan int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ch1 // waiting for signal

	for i := 0; i <= num; i++ {
		fmt.Println(text, i)
	}

	ch2 <- 1
}

func Second(text string, ch2, ch3 chan int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ch2

	for i := 0; i <= num; i++ {
		fmt.Println(text, i)
	}

	ch3 <- 1
}

func Third(text string, ch3 chan int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ch3

	for i := 0; i <= num; i++ {
		fmt.Println(text, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	go First("First Go Routine", ch1, ch2, 5, &wg)
	go Second("Second Go Routine", ch2, ch3, 5, &wg)
	go Third("Third Go Routine", ch3, 5, &wg)

	ch1 <- 1

	wg.Wait()
}
