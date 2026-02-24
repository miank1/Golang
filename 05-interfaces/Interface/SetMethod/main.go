package main

import "fmt"

type salaryCalculator interface {
	calculateSalary() float64
	report()
}

type PermanentEmployee struct {
	id          int
	basicSalary float64
	commission  float64
}

type ContractEmployee struct {
	id          int
	basicSalary float64
}

func (p PermanentEmployee) calculateSalary() float64 {
	return p.basicSalary + (p.commission/100)*p.basicSalary
}

func (c ContractEmployee) calculateSalary() float64 {
	return c.basicSalary
}

func (p PermanentEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month \n", p.id, p.calculateSalary())
}

func (c ContractEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month \n", c.id, c.calculateSalary())
}

func main() {
	p1 := PermanentEmployee{id: 1, basicSalary: 2300, commission: 13}
	p2 := PermanentEmployee{id: 2, basicSalary: 1500, commission: 18}
	p3 := PermanentEmployee{id: 3, basicSalary: 2300, commission: 10}
	c1 := ContractEmployee{id: 4, basicSalary: 500}
	c2 := ContractEmployee{id: 5, basicSalary: 1100}
	c3 := ContractEmployee{id: 6, basicSalary: 700}

	employees := []salaryCalculator{p1, p2, p3, c1, c2, c3}

	var totalSalary float64

	for _, employee := range employees {
		totalSalary += employee.calculateSalary()
	}

	fmt.Printf("Company total salary is : USD %f", totalSalary)
}
