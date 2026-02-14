package main

import (
	"fmt"
	"sync"
)

type Map struct {
	// Add fields
	data map[string]string
	mu   sync.Mutex
}

func NewMap() *Map {
	return &Map{
		data: make(map[string]string),
	}
}

func (m *Map) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value

}
func (m *Map) Get(key string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.data[key]
}

func main() {
	m1 := NewMap()
	m1.Set("name", "John")

	fmt.Println("Name :", m1.Get("name"))

}
