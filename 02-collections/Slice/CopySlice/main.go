// To copy a slice without sharing the underlying array, you use the built-in copy() function.

// This creates a new backing array and copies the elements into it.

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}

	// new slice with the same length
	s2 := make([]int, len(s1))

	// Copy contents
	copy(s2, s1)

	s2[0] = 100

	fmt.Println(s1) // [1 2 3 4]
	fmt.Println(s2) // [100 2 3 4]

	src := []int{1, 2, 3}
	dst := make([]int, 1)

	n := copy(dst, src)

	fmt.Println(n, dst)

}
