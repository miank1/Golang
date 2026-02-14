package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Employee struct {
	Name   string
	Age    int
	Salary int
}

type Employee1 struct {
	Name    string
	Age     int
	Salary  int
	Address Address
}

type Address struct {
	city    string
	country string
}

func main() {
	e := Employee{}
	fmt.Println(e)
	emp := Employee{
		Name:   "Sam",
		Age:    31,
		Salary: 2000,
	}

	fmt.Printf("emp : %v\n", emp)
	fmt.Printf("emp : %+v\n", emp)
	fmt.Printf("emp : %#v\n", emp)
	fmt.Println(emp)

	n := emp.Name
	fmt.Printf("Current Name is  : %+v\n", n)

	emp.Name = "John"
	fmt.Printf("New Name is : %s\n", emp.Name)

	// Creating struct without field names
	emp1 := Employee{"John", 32, 4000}

	fmt.Printf("emp : %v\n", emp1)
	fmt.Printf("emp : %+v\n", emp1)
	fmt.Printf("emp : %#v\n", emp1)
	fmt.Println(emp)

	// Pointer to struct
	empP := &emp1
	fmt.Printf("EmpP: %v\n", empP)

	// Another way of creating pointer to struct
	empP1 := &Employee{"John", 32, 4000}
	fmt.Printf("EmpP1: %v\n", empP1)

	// Marshall
	empJSON, err := json.Marshal(emp)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("MarshalIndent funnction output %s\n", string(empJSON))

	// Anonymous Fields
	type animal struct {
		string
		age int
	}

	a := animal{string: "Lion", age: 24}
	n = a.string

	fmt.Printf("Current name is: %s\n", n)
	a.string = "Tiger"
	fmt.Printf("New name is: %s\n", a.string)

	// Nested Struct

	addr := Address{city: "London", country: "UK"}

	emp4 := Employee1{Name: "Sam", Age: 31, Salary: 6000, Address: addr}

	fmt.Printf("Emp4 with address %v", emp4)

}
