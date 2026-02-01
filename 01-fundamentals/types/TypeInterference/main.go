package main

import "fmt"

func main() {
	var name = "John Mathew"
	fmt.Printf("Variable 'name' is of type %T \n", name)

	// This will throw a compiler error.
	// name = 1234

	var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0

	fmt.Printf("firstName : %T, lastName: %T, age : %T, salary: %T \n", firstName, lastName, age, salary)

	v := 42
	fmt.Printf("v is of type %T\n", v)

}
