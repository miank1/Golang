package main

import (
	"fmt"
	"log"
)

type employee struct {
	Name string
	Age  int
}

func main() {
	name := "John"
	age := 21
	fmt.Print("Name is : ", name, "\n")
	fmt.Print("Age is : ", age, "\n")

	e := employee{
		Name: name,
		Age:  age,
	}

	fmt.Print("Employee Struct ", e, "\n")

	// Number of Bytes Printed
	bytesPrinted, err := fmt.Print("Employee Struct ", e, "\n")

	if err != nil {
		log.Fatalln("Error Occured ", err)
	}

	fmt.Print(bytesPrinted)

	// Printing a string Variable and Integer variable
	fmt.Printf("Name is %s \n", name)
	fmt.Printf("Age is %d \n", age)

	// Printing struct
	fmt.Printf("Employee is %v \n", e)
	fmt.Printf("Employee is %+v \n", e)
	fmt.Printf("Employee is %#v \n", e)

}
