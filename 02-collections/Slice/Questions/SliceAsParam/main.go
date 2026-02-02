package main

import "fmt"

// modify the element
func modifySlice(slice []int) {
	slice = append(slice, 10202020)
	fmt.Println("Slice inside the function", slice)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice: ", slice)

	// Passing the slice to a function
	modifySlice(slice)
	fmt.Println("Final Slice:", slice)
}
