package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}

	fmt.Println(s1)
	fmt.Println(s2)

	Change(s1, s2)

	fmt.Println(s1)
	fmt.Println(s2)

}

func Change(s1, s2 []int) {

	s1 = append(s1, 6*2)

	fmt.Println("After append", s1)

	for i := 1; i < 5; i++ {

		s1[i] = s1[i] * i
		s2[i] = s2[i] * i

	}

	fmt.Println(s1)
	fmt.Println(s2)

}
