package main

import "fmt"

// Pass By Value /Reference using struct
type Person struct {
	Name string
	Age  int
}

func (p Person) UpdateAgeByValue(newAge int) {
	p.Age = newAge
}

func (p *Person) UpdateAgeByReference(newAge int) {
	p.Age = newAge
}

// pass by value with data type
func swap(a int, b int) {
	a, b = b, a
}

func swapRef(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 1, 2
	fmt.Println("Before swap ", a, b)
	swap(a, b)
	fmt.Println("After swap", a, b)
	c, d := 3, 4
	fmt.Println("Before swap ", c, d)
	swapRef(&c, &d)
	fmt.Println("After swap", c, d)

	p := Person{Name: "John", Age: 30}

	p.UpdateAgeByValue(40)
	fmt.Println(p.Age)

	p.UpdateAgeByReference(50)
	fmt.Println(p.Age)
}
