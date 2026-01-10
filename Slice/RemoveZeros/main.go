// Remove Zeros from a slice

package main

import "fmt"

func removeZeros(nums []int) []int {
	j := 0
	for _, v := range nums {
		if v != 0 {
			nums[j] = v
			j++
		}
		fmt.Println(nums)
	}
	return nums[:j]
}

func main() {
	s := []int{0, 1, 0, 3, 12}

	s = removeZeros(s)

	fmt.Println(s)
}
