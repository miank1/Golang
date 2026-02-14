package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Allocate some memory for the program to use
	s := make([]string, 0, 100000)
	for i := 0; i < 100000; i++ {
		s = append(s, "hello, world")
	}

	// Print the initial memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Initial HeapAlloc: ", m.HeapAlloc)

	// Trigger the garbage collector
	runtime.GC()

	// Print the memory usage after the garbage collector has run
	runtime.ReadMemStats(&m)
	fmt.Println("After GC HeapAlloc: ", m.HeapAlloc)

	// Release the memory
	s = nil
	// Trigger the garbage collector
	runtime.GC()
	// Print the memory usage after the garbage collector has run
	runtime.ReadMemStats(&m)
	fmt.Println("After release HeapAlloc: ", m.HeapAlloc)
}
