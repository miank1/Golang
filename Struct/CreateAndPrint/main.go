package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name string
	Age  int
	Address
}

func (p *Person) Greet() {
	fmt.Printf("Hello, my name is  %s and I am %d years old \n", p.Name, p.Age)
}

// Pointer receiver to modify Age
func (p *Person) HaveBirthday() {
	p.Age++ // increase age by 1
}

func main() {
	p := Person{
		Name: "Alen",
		Age:  24,
		Address: Address{
			City:    "Newyork",
			Country: "USA",
		},
	}

	fmt.Println(p)
	fmt.Println("Name and Age and City ->", p.Name, p.Age, p.City)

	p.Age = 25
	fmt.Println("Name and Age ->", p.Name, p.Age)

	ptr := &p

	ptr.Name = "Alice"

	fmt.Println(p)
	fmt.Println(ptr.Name)

	p.Greet()
	p.HaveBirthday() // increments age
	p.Greet()
}
