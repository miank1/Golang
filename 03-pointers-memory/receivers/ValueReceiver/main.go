package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type Lion struct {
	age int
}

type Cat struct {
	age int
}

func (l Lion) breathe() {
	fmt.Println("Lion Breathes")
}

func (l Lion) walk() {
	fmt.Println("Lion Walks")
}

func (c *Cat) breathe() {
	fmt.Println("Cat Breathes")
}

func (c *Cat) walk() {
	fmt.Println("Cat Walks")
}

func main() {
	var a animal

	a = Lion{age: 5}
	a.breathe()
	c1 := &Cat{age: 5}
	c1.walk()
	c2 := Cat{age: 5}
	c2.breathe()
}
