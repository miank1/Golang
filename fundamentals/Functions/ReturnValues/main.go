package main

import "fmt"

func add(a int, b int) (int, int) {
	return a, b
}

func main() {
	a, b := add(2, 4)
	fmt.Println(a, b)
	fmt.Printf("Total of a and b is %d \n", a+b)

	// ignoring the first value.
	_, c := add(2, 4)
	fmt.Println(c)
}
