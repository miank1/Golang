package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

type CirclePtr struct {
	radius float64
}

type RectanglePtr struct {
	length, width float64
}

func (c *CirclePtr) Area() float64 { // pointer receiver
	return 3.14 * c.radius * c.radius
}

func (r *RectanglePtr) Area() float64 { // pointer receiver
	return r.length * r.width
}

func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func main() {
	// You can create instances of Circle and Rectangle here and call TotalArea

	var shapes []Shape
	circle := Circle{radius: 5}
	rectangle := Rectangle{length: 2, width: 4}

	shapes = append(shapes, circle)
	shapes = append(shapes, rectangle)

	// Using pointer receivers
	shapes2 := []Shape{
		&CirclePtr{5},
		&RectanglePtr{2, 4}}

	fmt.Println("Total Area with pointer receivers:", TotalArea(shapes2))

	fmt.Println("Total Area :", TotalArea(shapes))

}
