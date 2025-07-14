package main

import "fmt"

func main() {
	var a = 5
	var b = 10
	swap(&a, &b)
	fmt.Println("After swap: a =", a, ", b =", b)
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
