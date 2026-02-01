package main

import (
	"fmt"
	hello "hello/PackageImport"
)

func main() {
	p := hello.Person{
		Name: "John",
	}
	fmt.Println("Accessing struct from another package - Person Name ->", p.Name)
	hello.SayHello()
}
