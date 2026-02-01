package main

import "fmt"

func main() {
	var floatValue float32 = 9.8

	// Conversion from floatValue to int Value
	var intValue int = int(floatValue)

	fmt.Println("The float value is %g", floatValue)
	fmt.Println("The int value is %d", intValue)

	// Implicit Casting - will not worl
	// var number int = 4.34

	var a float64 = 12
	var b int = int(a)

	fmt.Printf("Underlying Type of b: %T \n", b)
	b2 := int(a)
	fmt.Printf("Underlying Type of a: %T \n", b2)

}
