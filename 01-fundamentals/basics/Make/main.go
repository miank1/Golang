package main

import "fmt"

// In Go, make is a built-in function used to create slices, maps, and channels.
// It initializes and allocates memory for these composite data types and returns an initialized instance.
// In contrast, new([]int) returns a pointer to a newly allocated, zeroed slice structure, that is, a pointer to a nil slice value.

func main() {
	ss := make([]int, 5)
	fmt.Println("Slice", ss)
	fmt.Println("Length", len(ss))
	fmt.Println("Capacity", cap(ss))

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println("Map:", m)
	fmt.Println("Value for 'one':", m["one"])
	fmt.Println("Value for 'two':", m["two"])

	ch := make(chan int)
	go func() {
		ch <- 42
	}()

	value := <-ch
	fmt.Println("Received value from a channel ", value)
}
