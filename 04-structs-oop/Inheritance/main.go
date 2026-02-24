package main

import "fmt"

type iBase interface {
	say()
}

type base struct {
	color string
}

type child struct {
	base // base struct is embedded into child struct directly
	// with this child struct is able to access base method and data directly
	style string
}

func (b *base) say() {
	b.clear()
}

func (b *base) clear() {
	fmt.Println("Clear from base's function")
}

func check(b iBase) {
	b.say()
}

func main() {
	base := base{
		color: "Red",
	}

	child := &child{
		base:  base,
		style: "somestyle",
	}

	child.say()                               // calling base method
	fmt.Println("The color is ", child.color) // accessing base property
	check(child)

}
