package main

import "fmt"

type iBase interface {
	say()
}

type base struct {
	value string
}

type child struct {
	base
	color string
}

func (b *base) say() {
	fmt.Println(b.value)
}

func check(b iBase) {
	b.say()
}

func main() {
	base := base{
		value: "somevalue",
	}

	child := &child{
		base:  base,
		color: "Green",
	}
	child.say()
	check(child)
}
