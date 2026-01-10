package main

import "fmt"

func moveZeros(nums []int) {

	pos := 0

	for _, x := range nums {
		if x != 0 {
			nums[pos] = x
			pos++
		}
	}
	for i := pos; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}

func main() {

	ss := []int{0, 1, 0, 3, 12}
	moveZeros(ss)
}
