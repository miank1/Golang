package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, jobs chan int, wg *sync.WaitGroup) {
	for x := range jobs {
		fmt.Printf("worker %d processed job %d\n", i, x)
		time.Sleep(1 * time.Second)
		wg.Done()
	}

}

func producer(jobs chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		wg.Add(1) // One job added
		jobs <- i
	}
	close(jobs)
}

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	go producer(jobs, &wg)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, &wg)
	}

	wg.Wait()

}
