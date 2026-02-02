package main

// Sum of elements in a slice using go routines

import (
	"fmt"
	"sync"
)

func Sum(start, end int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}

	ch <- total

}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	l := len(arr)

	go Sum(0, l/2, &wg, ch)
	go Sum(l/2+1, l, &wg, ch)

	wg.Wait()
	close(ch)

	sum := 0

	for x := range ch {
		sum += x
	}
	fmt.Println("Sum of numbers is ", sum)

}
