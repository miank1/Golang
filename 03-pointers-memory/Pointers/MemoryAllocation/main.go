package main

import "fmt"

func main() {

	p := new(int)
	fmt.Println("Memory created ", p)
	*p = 10
	fmt.Println(*p, &p)

	s := make([]int, 3)
	s[0] = 10
	s[1] = 20
	s[2] = 30

	fmt.Println("Slice ", s)

}
