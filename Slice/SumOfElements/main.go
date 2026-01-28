package main

import "fmt"

// Write a function that takes a slice of integers as input and returns the sum of all the elements.

func SumOfElements(ss []int) {
	sum := 0
	for i := 0; i < len(ss); i++ {
		sum += ss[i]
	}

	fmt.Println(sum)
}

func main() {
	ss := []int{1, 2, 3, 4, 5}
	SumOfElements(ss)
}
