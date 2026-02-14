package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)

}
