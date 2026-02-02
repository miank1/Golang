package main

import "fmt"

func main() {

	a := make([]int, 5, 10)

	b := a[1:3]

	fmt.Println(len(b), cap(b)) // capacity of cap(a) - low of b = 10 - 1

	a1 := []int{1, 2, 3, 4, 5}
	b1 := a1[2:] // cap = 3 (5 - 2)

	fmt.Println(len(b1), cap(b1))

}
