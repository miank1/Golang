package main

import "fmt"

func modify(s []int) {
	s = append(s, 100)
	fmt.Println(s)
}

func main() {

	arr := []int{1, 2, 3}
	modify(arr)

	fmt.Println(arr)

}
