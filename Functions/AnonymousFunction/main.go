package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World !!!")
	}()

	res := max(2, 3)
	fmt.Println(res)
}

var max = func(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
