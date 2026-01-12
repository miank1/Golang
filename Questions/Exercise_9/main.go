// WaitGroup + Goroutines

package main

import (
	"fmt"
	"sync"
)

func Print(str string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println(str)
	}

}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go Print("A", &wg)
	go Print("B", &wg)
	go Print("C", &wg)

	wg.Wait()
}
