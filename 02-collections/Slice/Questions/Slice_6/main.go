package main

import "fmt"

// Slice pass by value and reference

func PassByValue(ss []int) {
	ss[0] = 10
	fmt.Println(ss)
}

func main() {
	ss := []int{1, 2, 3, 4, 5}
	PassByValue(ss)
	fmt.Println(ss)
}
