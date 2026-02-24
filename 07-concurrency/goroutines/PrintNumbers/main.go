// You are given a function called printNumbers that prints numbers from 1 to 5. Implement a program that uses Go routines to execute printNumbers concurrently,
// ensuring that the output is in the sequence: 12345.

package main

import (
	"fmt"
	"sync"
)

func PrintNumbers(wg *sync.WaitGroup, num []int, ch chan int) {
	defer wg.Done()

	<-ch

	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}

	ch <- 1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int, 1)

	ch <- 1
	go PrintNumbers(&wg, arr, ch)
	wg.Wait()

}
