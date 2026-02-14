package model

import "fmt"

var (
	CompanyName     = "test"     // Exported
	companyLocation = "somecity" // Not exported
)

// Person Struct

type Person struct {
	Name string
	age  int
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) getName() string {
	return p.Name
}

type company struct {
}

func GetPerson() *Person {
	p := &Person{
		Name: "test",
		age:  21,
	}

	fmt.Println("Model Package:")
	fmt.Println(p.Name)
	fmt.Println(p.age)
	return p
}

func getCompnayName() string {
	return CompanyName
}
