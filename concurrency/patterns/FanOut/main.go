package main

import (
	"fmt"
	"time"
)

func worker(id int, task int) {
	fmt.Printf("Worker %d is processing task %d \n", id, task)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d has finished task %d \n", id, task)
}

func main() {

	tasks := []int{1, 2, 3}

	for i, task := range tasks {
		go worker(i, task)
	}

	time.Sleep(2 * time.Second)

}
