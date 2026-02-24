// Maps are reference type in go
// When you pass a map to a function the function sees the underlying map.

package main

import "fmt"

func add(m map[string]int) {
	m["a"] = 1
}

func edge_case(m map[string]int) {
	m = make(map[string]int)
	m["a"] = 1
}

func main() {
	m := make(map[string]int)
	add(m)
	fmt.Println(m["a"]) // 1
	m1 := make(map[string]int)
	edge_case(m1)
	fmt.Println(m1["a"])
}
