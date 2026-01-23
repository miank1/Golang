// goroutine calculates square of a number

package main

import (
	"fmt"
	"sync"
)

func printSquare(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i * i)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go printSquare(i, &wg)
	}

	wg.Wait()
}
