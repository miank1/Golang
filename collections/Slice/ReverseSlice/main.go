// Reverse a slice
// Input:  []int{1,2,3,4}
// Output: []int{4,3,2,1}

package main

import "fmt"

func reverse(nums []int) {

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

}
func main() {
	s := []int{1, 2, 3, 4}
	reverse(s)
	fmt.Println(s)
}
