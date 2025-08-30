package main

import (
	"fmt"
	"sort"
)

func main() {

	m := map[string]int{
		"banana": 3,
		"apple":  5,
		"cherry": 2,
	}

	// Step 1: Extract keys
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Step 2: Sort keys
	sort.Strings(keys)

	// Step 3: Iterate in sorted order
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, m[k])
	}

}
