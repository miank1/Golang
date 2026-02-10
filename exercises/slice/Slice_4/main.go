package main

import "fmt"

// Write a program that finds the intersection of two slices.
// Given two slices, return a new slice that contains only the common elements.

func IntersectionOfSlice(ss1, ss2 []int) {
	m1 := make(map[int]int)

	for i := 0; i < len(ss1); i++ {
		m1[ss1[i]]++
	}

	for i := 0; i < len(ss2); i++ {
		m1[ss2[i]]++
	}

	for x, v := range m1 {
		if v > 1 {
			fmt.Println("Common Elements are ", x)
		}
	}
}

func main() {
	ss1 := []int{1, 2, 3, 4, 7}
	ss2 := []int{2, 2, 4, 4, 7}

	IntersectionOfSlice(ss1, ss2)
}
