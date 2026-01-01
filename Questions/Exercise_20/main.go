// Fan-Out + Fan-In (Parallel Workers + Merge)
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(300 * time.Millisecond)
		result := job * job
		fmt.Printf("Worker %d processed job %d -> %d\n", id, job, result)
		results <- result
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	out := make(chan int)

	// Start 3 workers (fan-out)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Send 5 jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// Fan-in goroutine: collect results
	go func() {
		for i := 1; i <= 5; i++ {
			out <- <-results
		}
		close(out)
	}()

	// Print merged results
	for x := range out {
		fmt.Println("Output:", x)
	}
}
