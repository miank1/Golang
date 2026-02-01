// Input:  [7,1,5,3,6,4]
// input - prices[i], the price of an item on day i.
// You may buy once and sell once later.

// return maximum profit

// [9,9,1,1,7,7,4]

package main

import "fmt"

func singleElement_1(arr []int) int {

	ans := 0

	for _, v := range arr {
		ans ^= v
	}

	return ans
}

func singleElement(arr []int) int {

	m1 := make(map[int]int)

	for _, v := range arr {
		m1[v]++
	}

	for x, v := range m1 {
		if v == 1 {
			return x
		}
	}

	return -1
}

func main() {

	arr := []int{9, 9, 1, 1, 7, 7, 4}

	result := singleElement(arr)
	result1 := singleElement_1(arr)
	fmt.Println("Single element is  ", result)
	fmt.Println("Single element is  ", result1)
}
