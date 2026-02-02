package main

import "fmt"

type Student struct {
	Name  string
	Marks int
}

func updateMarks(registry map[string]*Student, name string, newMarks int) {

	if students, exists := registry[name]; exists {
		students.Marks = newMarks
	}

}

func addStudent(registry map[string]*Student, name string, marks int) {
	if _, exists := registry[name]; exists {
		fmt.Println("student already exists :", name)
		return
	}

	registry[name] = &Student{Name: name, Marks: marks}

}

func main() {

	students := make(map[string]*Student)

	// Adding student
	addStudent(students, "Alice", 85)
	addStudent(students, "Bob", 90)
	addStudent(students, "Alice", 70)

	updateMarks(students, "Alice", 92)

	fmt.Println(*students["Alice"], *students["Bob"])
}
