package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) details() {
	fmt.Printf("Name: %s\n", e.name) // string
	fmt.Printf("Age: %d\n", e.age)   // age
}

func (e employee) getSalary() int {
	return e.salary
}

func (e employee) setNewName(newName string) {
	e.name = newName
}

// Method on a Pointer Receiver

func (e *employee) setNewName1(newName string) {
	e.name = newName
}

func main() {
	emp := employee{
		name:   "Sam",
		age:    31,
		salary: 2000,
	}
	emp.details()
	fmt.Printf("Salary %d \n", emp.getSalary())
	emp.setNewName("John")
	fmt.Printf("Name :%s \n", emp.name)

	emp1 := &employee{
		name: "Sam", age: 31, salary: 2000}

	emp1.setNewName1("John")
	fmt.Printf("Name : %s\n", emp1.name)
}
