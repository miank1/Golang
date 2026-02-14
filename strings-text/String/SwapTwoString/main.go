package main

import "fmt"

func main(){
	a := "123"
	b := "456"

	fmt.Println("a ->", a)
	fmt.Println("b ->", b)
	a,b = b, a
	fmt.Println("a ->", a)
	fmt.Println("b ->", b)
}