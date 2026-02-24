package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")

	sem := make(chan struct{}, 3)

	for i := 1; i < 10; i++ {
		sem <- struct{}{}

		go func(id int) {
			fmt.Printf("Student %d is using the computer \n", id)
			time.Sleep(2 * time.Second)
			fmt.Printf("Studnet %d is done\n", id)
			<-sem
		}(i)
	}

	time.Sleep(5 * time.Second)
}
