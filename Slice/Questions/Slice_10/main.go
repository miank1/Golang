// Double each element of the slice

package main

import "fmt"

func doubleElements(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * 2
	}

	return nums

}

func main() {

	s1 := []int{1, 2, 3, 4, 5, 6}

	s1 = doubleElements(s1)

	fmt.Println(s1)
}
