package main

import "fmt"

func doubleElements(s *[]int) {

	for i, v := range *s {
		(*s)[i] = v * 2
	}

}

func main() {
	s := []int{1, 2, 3}
	doubleElements(&s)
	fmt.Println("After doubling:", s)
}
