package main

import (
	"fmt"
	"strconv"
)

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name string
	Age  int
	Address
}

type Employee struct {
	Person
	Position string
}

func (p *Person) Greet() {
	fmt.Printf("Hello, my name is  %s and I am %d years old \n", p.Name, p.Age)
}

// Pointer receiver to modify Age
func (p *Person) HaveBirthday() {
	p.Age++ // increase age by 1
}

func (a *Address) Greet() {
	fmt.Printf("I live in %s, %s", a.City, a.Country)
}

type Describable interface {
	Describe() string
}

func (a *Address) Describe() string {
	return "Address: " + a.City + "," + a.Country
}

func (p *Person) Describe() string {
	return "Person: " + p.Name + "," + strconv.Itoa(p.Age)
}

func (e *Employee) Describe() string {
	return "Employee: " + e.Person.Name + "," + "Age: " + strconv.Itoa(e.Person.Age) + "," + "Position:" + e.Position + "," + "City: " + e.Person.City
}

func main() {
	// Create instances
	p := &Person{
		Name: "Alice",
		Age:  25,
		Address: Address{
			City:    "Newyork",
			Country: "USA",
		},
	}

	a := &Address{City: "Berlin", Country: "Germany"}

	e := &Employee{
		Person: Person{
			Name: "Bob",
			Age:  30,
			Address: Address{
				City:    "London",
				Country: "UK",
			},
		},
		Position: "Engineer",
	}

	// Polymorphism in action
	things := []Describable{p, a, e}

	for _, t := range things {
		fmt.Println(t.Describe())
	}
}
