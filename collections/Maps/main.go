package main

import "fmt"

func main() {

	// Nil map read is allowed and write will cause panic.

	var m map[string]int
	fmt.Println("m ", m == nil) // true

	// Empty Map - read and write both allowed not a nil map
	m1 := make(map[string]int)
	fmt.Println("m1 ", m1 == nil) // false
	fmt.Println("Length of empty map", len(m1))

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

	// delete from map
	m5 := map[string]int{
		"a": 1,
		"b": 2,
	}

	delete(m, "a")
	fmt.Println(m5) // go gurantees safe delete

}
