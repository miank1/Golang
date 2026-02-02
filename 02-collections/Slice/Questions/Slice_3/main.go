package main

import "fmt"

// Write a program that removes the duplicates from a given slice of integers. The resulting slice should only contain unique elements.

func RemoveDuplicates(ss []int) {
	m1 := map[int]int{}
	ss1 := make([]int, 0)

	for i := 0; i < len(ss); i++ {
		m1[ss[i]]++
	}

	for x, v := range m1 {
		fmt.Println(x, "->", v)
		ss1 = append(ss1, x)
	}

	fmt.Println(ss1)
}

func main() {
	ss := []int{1, 2, 2, 3, 4, 5, 6, 5}
	RemoveDuplicates(ss)
}
