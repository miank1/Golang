package main

import "fmt"

func main() {
	a := 2
	b := &a

	fmt.Println(a)
	fmt.Println(*b)

	*b = 3
	fmt.Println(a)
	fmt.Println(*b)

	a = 4
	fmt.Println(a)
	fmt.Println(*b)
}
