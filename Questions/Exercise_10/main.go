// Worker Pool

package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		result := job * 2
		fmt.Printf("Worker %d processed job %d -> result %d\n", id, job, result)
		results <- result
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}

}
