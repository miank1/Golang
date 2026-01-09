package main

import "fmt"

func sumSlice(nums []int) int {

	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}

func main() {
	ss := []int{1, 2, 3, 4}

	result := sumSlice(ss)

	fmt.Println(result)
}
