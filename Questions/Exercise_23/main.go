package main

import "fmt"

func main() {
	a := make([]int, 0, 2)
	a = append(a, 1, 2)

	b := a
	b = append(b, 3)

	fmt.Println(a)
	fmt.Println(b)
}
