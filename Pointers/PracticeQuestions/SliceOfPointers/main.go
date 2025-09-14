package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func main() {

	employee := []*Employee{
		&Employee{Name: "Alice", Salary: 1000},
	}

	fmt.Println("Employee ", employee) 

}
