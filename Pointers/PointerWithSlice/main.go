package main

import "fmt"

func appendNumber(s *[]int, n int) {

	*s = append(*s, 4, 5)
}

func main() {
	ss := []int{1, 2, 3}

	fmt.Println("Before Updating the slice", ss)
	appendNumber(&ss, len(ss))
	fmt.Println("After Updating the slice", ss)
}
