package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		if x < 10 {
			fmt.Println(x)
		}
		x++
	}

	if x%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
