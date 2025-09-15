package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, jobs chan int, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()
	for x := range jobs {
		result <- x * x
		fmt.Printf("worker %d processed job %d\n", i, x)
		time.Sleep(1 * time.Second)
	}
}

func producer(jobs chan int) {
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	result := make(chan int, 10)
	go producer(jobs)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg, result)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println("Result:", res)
	}

}
