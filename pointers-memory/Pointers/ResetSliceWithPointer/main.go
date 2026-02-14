package main

import "fmt"

func resetSlice(s *[]int) {

	for i := range *s {
		(*s)[i] = 0
	}
}

func main() {
	nums := []int{10, 20, 30, 40}
	fmt.Println("Before Slice Reset", nums)
	resetSlice(&nums)
	fmt.Println("After Slice Reset", nums)

}
