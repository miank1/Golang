package main

import "fmt"

func main() {
	a := []int{1, 2, 3} // length = 3 and capacity = 4
	b := a              // both shares same array

	a = append(a, 4) // new array allocated - capacity exceeded

	fmt.Println("b :", b) // b points to old array
	fmt.Println("a :", a) // a points to new array

	a1 := make([]int, 0, 3)
	b1 := a1

	a = append(a1, 1)
	a = append(a1, 2)
	a = append(a1, 3)
	a = append(a1, 4)

	fmt.Println("a1", a1)
	fmt.Println("b1", b1)

}
