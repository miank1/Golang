package main

import (
	"fmt"
	"sync"
)

func calculateSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := num * num
	fmt.Printf("%d squared is %d \n", num, square)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, &wg)
	}
	wg.Wait()
}
