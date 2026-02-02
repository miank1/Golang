// Concurrency Pattern using directional channels in a pipeline pattern
// Producer -> Processor -> Consumer

package main

import (
	"fmt"
	"sync"
)

func generator(nums chan<- int) {
	for i := 1; i <= 5; i++ {
		nums <- i
	}
	close(nums)
}

func square(nums <-chan int, output chan<- int) {
	for val := range nums {
		output <- val * val
	}
	close(output)
}

func printer(in <-chan int) {
	for result := range in {
		fmt.Println("Results :", result)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	numbers := make(chan int)
	squared := make(chan int)

	go generator(numbers)
	go square(numbers, squared)
	printer(squared)

}
