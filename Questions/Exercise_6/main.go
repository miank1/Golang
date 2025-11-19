// Interfaces & Polymorphism

package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return c.Radius * 3.14
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	circle := Circle{
		Radius: 5,
	}

	rectangle := Rectangle{
		Width:  4,
		Height: 6,
	}

	PrintArea(circle)
	PrintArea(rectangle)

}
