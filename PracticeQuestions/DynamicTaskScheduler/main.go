// Problem: Dynamic Task Scheduler with Interface-based Task Execution
// Problem Statement:
// You are building a dynamic task scheduler that must execute different types of tasks concurrently. Each task implements a common interface called Task, which defines a Run() method.

// Design a system with the following requirements:

// Define a Task interface with a Run() error method.

// Create at least two concrete types that implement the Task interface, e.g., EmailTask, ReportTask.

// The Scheduler should:

// Accept any number of tasks (implementing the Task interface).

// Run N tasks concurrently (limit concurrency).

// Collect and print the result (success or error) of each task.

// Print the output in the same order the tasks were submitted (even though they run in parallel).

package main

import "fmt"

type Task interface {
	Run() error
}

type EmailTask struct {
}

type ReportTask struct {
}

func (e EmailTask) Run() error {
	fmt.Println("Sending email ...")
	return nil
}

func (r ReportTask) Run() error {
	fmt.Println("Generating Report ...")
	return nil
}

func Scheduler(s []int, concurrency int) {

	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
}

func main() {

}
