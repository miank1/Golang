package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func updateAge(p *Person, newAge int) {
	p.Age = newAge
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  25,
	}

	fmt.Println("Person Details ", p)
	updateAge(&p, 26)
	fmt.Println("Person Details ", p)

}
