package main

import "fmt"

func removeAt(nums []int, idx int) []int {

	if idx < 0 || idx >= len(nums) {
		return nums
	}
	return append(nums[:idx], nums[idx+1:]...)

}

func main() {
	nums := []int{10, 20, 30, 40}
	fmt.Println(removeAt(nums, 2))
}
