package main

import (
	"fmt"
	"sync"
)

func sumOfSlice(slice []int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	sum := 0

	for _, num := range slice {
		sum += num
	}

	resultChan <- sum

}

func main() {

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceSize := len(num)

	var wg sync.WaitGroup

	resultChan := make(chan int, sliceSize)

	mid := sliceSize / 2

	firstHalf := num[:mid]
	secondHalf := num[mid:]

	wg.Add(2)

	go sumOfSlice(firstHalf, &wg, resultChan)
	go sumOfSlice(secondHalf, &wg, resultChan)

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	fmt.Println("Sum of slice is :", totalSum)

}
