package main

import "fmt"

type Student struct {
	Name  string
	Marks int
}

func updateMarks(m map[string]int, name string, marks int) {

	m[name] = 98

}

func main() {

	students := make(map[string]int)

	students["Alice"] = 78
	fmt.Println("Before Update ", students)

	updateMarks(students, "Alice", 98)
	fmt.Println("After Update ", students)

}
