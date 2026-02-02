// Print Even and Odd Numbers Using Goroutines
// Create two goroutines: one prints even numbers, the other odd numbers up to n, taking turns.

package main

import (
	"fmt"
	"sync"
)

func Even(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i += 2 {
		<-ch // wait for signal
		fmt.Println("Even:", i)
		ch <- 1 // signal odd
	}
}

func Odd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-ch // wait for signal
		fmt.Println("Odd:", i)
		ch <- 0 // signal even
	}

}

func main() {

	ch := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go Even(ch, &wg)
	go Odd(ch, &wg)
	ch <- 1

	wg.Wait()
	close(ch)
}
