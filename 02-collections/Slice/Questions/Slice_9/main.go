package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// convert to slice
	s1 := arr[:]
	fmt.Println("Array :", arr)
	fmt.Println("Slice :", s1)

	s1 = append(s1, 4)

	fmt.Println("Updated slice :", s1)

}
