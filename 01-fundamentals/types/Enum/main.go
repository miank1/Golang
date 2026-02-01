package main

import "fmt"

const (
	smallSize = 2
	medium    = iota
	large
	extraSize
)

func main() {
	fmt.Println(smallSize)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraSize)
}
