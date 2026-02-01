// constant in golang
package main

import "fmt"

const (
	r = "Rectange"
	s = "Square"
	x = 242
)

func main() {
	const c string = "circle"
	fmt.Println(c)
	const d = 8
	fmt.Println(d)
	fmt.Println(r, " ", s)
	// const b = getValue() // not allowed as const value must be known at compile time.
	// fmt.Println(b)
	const x = 456
	fmt.Println(x)
}

func getValue() int {
	return 1
}
