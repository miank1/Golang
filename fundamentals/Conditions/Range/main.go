package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}

	for i, v := range nums {
		fmt.Println(i, v)
	}

	m := map[string]int{"a": 1, "b": 2}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
