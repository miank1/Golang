package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, n := range nums {
		total += n
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(5, 10))
	fmt.Println(sum())

}
