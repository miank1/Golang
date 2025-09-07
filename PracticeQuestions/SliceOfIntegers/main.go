package main

import "fmt"

func EvenNumbers(nums []int) (evenNumbers []int) {

	if nums == nil {
		return []int{}
	}

	evenNum := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			evenNum = append(evenNum, nums[i])
		}
	}

	return evenNum
}

func main() {

	fmt.Println("Input:", []int{1, 2, 3, 4, 5, 6})
	fmt.Println("Output:", EvenNumbers([]int{1, 2, 3, 4, 5, 6}))

	fmt.Println("Input:", []int{})
	fmt.Println("Output:", EvenNumbers([]int{}))

	fmt.Println("Input: nil")
	fmt.Println("Output:", EvenNumbers(nil))

}
