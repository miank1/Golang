package main

import "fmt"

func main() {
	a := "123"
	b := "456"

	fmt.Printf("Before a:%s and b :%s \n ", a, b)
	a, b = b, a
	fmt.Printf("After a:%s and b :%s", a, b)
}
