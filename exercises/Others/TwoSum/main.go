package main

import "fmt"

func TwoSum(s []int, target int) []int {

	left, right := 0, len(s)-1

	for left < right {

		sum := s[left] + s[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(TwoSum(nums, target))

}
