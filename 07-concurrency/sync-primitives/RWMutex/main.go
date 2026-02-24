package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	m  map[string]string
	mu sync.RWMutex
}

func (s *SafeMap) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[key]
}

func (s *SafeMap) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func main() {
	s := SafeMap{m: make(map[string]string)}
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		s.Set("foo", "bar")
		wg.Done()
	}()

	// Multiple readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Printf("Reader %d read: %s\n", id, s.Get("foo"))
			wg.Done()
		}(i)
	}

	wg.Wait()

}
