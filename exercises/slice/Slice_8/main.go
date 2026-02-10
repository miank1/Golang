package main

import "fmt"

func main() {
	s1 := make([]int, 3, 5)
	s1[0], s1[1], s1[2] = 10, 20, 30

	fmt.Println(s1)

	s1 = append(s1, 40, 50)

	fmt.Println(s1)

}
