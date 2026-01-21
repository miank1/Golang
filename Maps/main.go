package main

import "fmt"

func main() {

	m1 := make(map[string]int) // m1 ready to use not a nil map
	fmt.Println("m1 ", m1 == nil)

	m2 := map[string]int{
		"apple":  10,
		"banana": 20,
	}

	// read from map

	fmt.Println("m2 :", m2["apple"])

	// iterate over the map

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// safe read

	v, ok := m2["apple"]

	if ok {
		fmt.Println("Apple Value", v)
	}

	// maps are Reference Types
	m3 := make(map[string]int)

	m4 := m3

	m4["a"] = 1
	fmt.Println(m3["a"]) // 1

}
