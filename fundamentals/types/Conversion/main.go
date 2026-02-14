package main

import "fmt"

func main() {
	var a float64 = 12
	var b int = int(a)

	fmt.Printf("Underlying type of b: %T\n", b)

	b2 := int(a)

	fmt.Printf("Underlying type of b: %T\n", b2)

	var c int = 12
	var d float64 = float64(c)

	fmt.Printf("Underlying Type of d: %T\n", d)

	d2 := float64(c)
	fmt.Printf("Underlying Type of d2: %T\n", d2)
}
