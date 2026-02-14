package main

import (
	"fmt"
	"time"
)

type Task struct {
	id int
}

func worker(id int, jobs <-chan Task, result chan<- string) {
	for task := range jobs {
		fmt.Printf("Worker  %d started task %d\n", id, task.id)
		time.Sleep(time.Second) // simulate time-consuming
		fmt.Printf("Worker %d finished task %d", id, task.id)
		result <- fmt.Sprintf("Task %d done by worker %d", task.id, id)
	}
}

func main() {
	numWorkers := 3
	numJobs := 5

	jobs := make(chan Task, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- Task{id: j}
	}
	close(jobs)

	for j := 1; j <= numJobs; j++ {
		fmt.Println(<-results)
	}

}
