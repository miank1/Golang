// Arrays, Slices & Range Loop

package main

import "fmt"

func numSlice(s []int) int {

	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(numSlice(nums))

}
