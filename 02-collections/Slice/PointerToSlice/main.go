package main

import "fmt"

func modifySlice(slice *[]int) {
	*slice = append(*slice, 10202020)
	fmt.Println("Slice inside the function:", slice)
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice :", slice1)

	// passing the slice to the function
	modifySlice(&slice1)
	fmt.Println("Final Slice: ", slice1)
}
