package main

import "fmt"

func mergeMaps(m1, m2 map[string]int) map[string]int {
	merged := make(map[string]int)

	for k, v := range m1 {
		merged[k] = v
	}

	for k, v := range m2 {
		merged[k] += v // this adds values from both maps
	}

	return merged
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}

	result := mergeMaps(m1, m2)
	fmt.Println(result)
}
