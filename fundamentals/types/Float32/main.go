package main

import "fmt"

func main() {
	var a float32 = 12.0
	var b float64 = float64(a)

	fmt.Printf("Underlying Type of b: %T\n", b)

	b2 := float64(a)
	fmt.Printf("Underlying Type of b2: %T\n", b2)

	var a1 float64 = 12.0
	var b1 float32 = float32(a1)

	fmt.Printf("Underlying Type of b: %T\n", b1)

	b1 = float32(a1)
	fmt.Printf("Underlying Type of b2 is  %T\n", b1)
}
