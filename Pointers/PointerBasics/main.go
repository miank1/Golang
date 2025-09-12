package main

import "fmt"

func main() {

	x := 10
	p := &x
	fmt.Println("Value of p and x ", *p, x)
	*p = 20
	fmt.Println("Value of p and x ", *p, x)

}
