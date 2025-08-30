package main

import "fmt"

type Animal struct{}

func (a Animal) Speak() {
	fmt.Println("Animal speaks")
}

type Dog struct {
	Animal
}

func (d Dog) Speak() {
	fmt.Println("Dog barks")
}

func makeItSpeak(a Animal) {
	a.Speak()
}

func main() {
	d := Dog{}
	d.Speak()
	d.Animal.Speak()
	makeItSpeak(d.Animal)
}
