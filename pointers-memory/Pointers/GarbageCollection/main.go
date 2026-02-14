package main

import (
	"fmt"
	"runtime"
)

func allocateSlice() {
	s := make([]int, 1000000)
	for x := range s {
		s[x] = x
	}
}

func main() {
	fmt.Println("Before allocation:")
	printMemStats()
	allocateSlice()
	fmt.Println("After allocation, before GC:")
	printMemStats()
	runtime.GC()
	fmt.Println("After GC:")
	printMemStats()
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
}
