package main

import (
	"fmt"
	"time"
)

func Task(name string) {
	fmt.Println("Running Task:", name, "at", time.Now())
}

func startScheduler(interval time.Duration, taskName string) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go Task(taskName)
		}
	}
}

func main() {
	go startScheduler(5*time.Second, "TaskA")
	go startScheduler(10*time.Second, "TaskB")

	select {}

}
