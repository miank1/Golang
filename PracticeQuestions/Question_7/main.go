package main

import "fmt"

func main() {
	ss := []int{1, 2, 3, 4, 5, 6}
	ss1 := []int{}

	for i := 0; i < len(ss); i++ {
		if ss[i]%2 != 0 {
			ss1 = append(ss1, ss[i])
		}
	}

	fmt.Println(ss1)

}
