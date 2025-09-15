// Write a Go program that launches a goroutine which prints numbers from 1 to 5.
// while the main function prints letters from A to E.

package main

import (
	"fmt"
	"sync"
)

func Print(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go Print(&wg)

	fmt.Println("A")
	fmt.Println("B")
	fmt.Println("C")
	fmt.Println("D")
	fmt.Println("E")

	wg.Wait()
}
