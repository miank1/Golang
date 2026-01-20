package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 2, 1, 5}

	// Remove duplicates from a slice

	m := make(map[int]bool)
	result := []int{}

	for _, v := range arr1 {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}

	fmt.Println(result)

}
