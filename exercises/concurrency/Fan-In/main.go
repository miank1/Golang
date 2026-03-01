package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	results := make(chan int, len(nums))

	var wg sync.WaitGroup

	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			results <- num * num
		}(n)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for val := range results {
		sum += val
	}

	fmt.Println(sum) // 55
}
