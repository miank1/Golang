package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
	handle(1, "abc")
	handle("abc", "xyz", 3)
	handle(1, 2, 3, 4)

}

// Using Vardiac function - Accepts variable number of arguments
func add(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return sum
}

// This can handle both varidiac function and empty interface
func handle(params ...interface{}) {
	fmt.Println("Handle func called with parameters")
	for _, param := range params {
		fmt.Printf("%v\n", param)
	}
}
