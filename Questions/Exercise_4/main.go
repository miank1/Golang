// Structs & Methods

package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {

	return r.Height * r.Width
}

func main() {
	rect := Rectangle{
		Width:  10,
		Height: 20,
	}
	fmt.Println(rect.Area())
}
