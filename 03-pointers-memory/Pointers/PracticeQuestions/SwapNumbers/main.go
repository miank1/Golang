package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 10, 20

	fmt.Println("Before Swap ", a, b)
	swap(&a, &b)
	fmt.Println("After Swap ", a, b)

}
