package main

import "fmt"

func main() {
	s := make([]int, 2, 5)
	s[0], s[1] = 1, 2

	t := append(s, 3)

	fmt.Println(s)
	fmt.Println(t)
}
