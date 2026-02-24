package main

import (
	"fmt"
	"sync"
	"time"
)

func serviceA(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Service A is done")
}

func serviceB(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Service B is done")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go serviceA(&wg)
	go serviceB(&wg)

	wg.Wait() // wait until counter is = 0

	fmt.Println("All done")
}
