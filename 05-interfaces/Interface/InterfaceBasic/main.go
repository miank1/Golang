// Interface define behaviour not the data

package main

import "fmt"

type Speaker interface {
	Speak() string
}

// Any type that implements menthods automatically satisifes the interface

type Human struct{}

func (h Human) Speak() string {
	return "Hello from Human"
}

type Robot struct{}

func (r Robot) Speak() string {
	return "Hello from Robot"
}

func saySomething(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {

	// using the interface

	h := Human{}
	saySomething(h)

	r := Robot{}
	saySomething(r)
}
