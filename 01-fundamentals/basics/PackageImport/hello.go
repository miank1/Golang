package hello

import "fmt"

type Person struct {
	Name string
}

func SayHello() {
	fmt.Println("Hello from anothe package")
}
