package main

import (
	"fmt"
	"sync"
)

type Counter interface {
	Inc()
	Value() int
}

type Count struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Count {
	return &Count{}
}

func (c *Count) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++

}

func (c *Count) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

func main() {
	c := NewCounter()

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()

	fmt.Println(c.Value()) // should print 1000
}
