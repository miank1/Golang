package main

import "fmt"

func add(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println(add(3, 3))
}
