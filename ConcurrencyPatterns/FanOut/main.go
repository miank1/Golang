package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d \n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d \n", id, job)
		results <- job * 2
	}
}

func main() {

	jobs := make(chan int, 10)
	results := make(chan int, 5)

	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	// collect results

	for i := 1; i <= 5; i++ {
		fmt.Println("Result:", <-results)
	}

}
