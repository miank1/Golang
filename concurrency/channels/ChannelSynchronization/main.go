package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working ...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	<-done
	fmt.Println("Main finished")
}
