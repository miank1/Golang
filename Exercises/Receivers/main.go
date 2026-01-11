package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Scale(factor float64) {
	r.width = r.width * factor
	r.height = r.height * factor

}

func main() {

	r := Rectangle{
		width:  4,
		height: 5,
	}

	fmt.Println("Before scaling:", r)

	r.Scale(2) // modifies original

	fmt.Println("After scaling by 2:", r)
	fmt.Println("Area:", r.Area())
}
