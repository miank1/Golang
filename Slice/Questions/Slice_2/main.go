package main

import "fmt"

// check if two slices are equal

func CompareSlices(ss1, ss2 []int) (str string, result bool) {

	if len(ss1) != len(ss2) {
		return "Both the slices are not equal", false
	}

	str1 := "Both the slices are equal"
	result1 := true
	for i := 0; i < len(ss1); i++ {
		if ss1[i] != ss2[i] {
			str1 = "Slices are not equal"
			result1 = false
			break
		}
	}

	return str1, result1
}

func main() {
	ss1 := []int{1, 2, 3, 4, 5}
	ss2 := []int{1, 2, 3, 4, 15}
	fmt.Println(CompareSlices(ss1, ss2))
}
