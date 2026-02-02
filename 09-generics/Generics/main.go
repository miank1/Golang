package main

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// Number is a constraint that allows int32 or float64.
type Number interface {
	int32 | float64
}

// Sum calculates the sum of a slice of values of type T (int32 or float64).
func Sum[T Number](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

func main() {
	ints := []int{1, 2, 3}
	strings := []string{"One", "Two", "Three"}
	Print(ints)
	Print(strings)

	// Example with int32
	int32Slice := []int32{1, 2, 3, 4, 5}
	int32Sum := Sum(int32Slice)
	fmt.Printf("Sum of int32 slice: %d\n", int32Sum)

	// Example with float64
	float64Slice := []float64{1.5, 2.5, 3.5}
	float64Sum := Sum(float64Slice)
	fmt.Printf("Sum of float64 slice: %.2f\n", float64Sum)
}
