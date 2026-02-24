package model

import (
	"fmt"
)

func Test() {
	p := &Person{
		Name: "test",
		age:  21,
	}

	fmt.Println(p)
	c := &company{}
	fmt.Println(c)

	// Structure Methods
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

	//STRUCTURE'S FIELDS
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//FUNCTION
	person2 := GetPerson()
	fmt.Println(person2)
	companyName := getCompnayName()
	fmt.Println(companyName)

	//VARIBLES
	fmt.Println(companyLocation)
	fmt.Println(CompanyName)
}
