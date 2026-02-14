package main

import "fmt"

func main() {

	x := 10

	switch {
	case x > 0:
		fmt.Println("positive")
	case x < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}

}
