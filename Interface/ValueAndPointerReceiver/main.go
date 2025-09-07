package main

import "fmt"

type Incrementer interface {
	Increment()
}

type CounterVal struct {
	Value int
}

func (c CounterVal) Increment() {
	c.Value++
}

type CounterPtr struct {
	Value int
}

func (c *CounterPtr) Increment() {
	c.Value++ // Actual object modification
}

func main() {

	// Value Receiver
	cVal := CounterVal{Value: 0}
	var incVal Incrementer = cVal

	incVal.Increment()
	fmt.Println("Value receiver : ", cVal.Value)

	// Pointer receiver
	cPtr := CounterPtr{Value: 0}
	var incPtr Incrementer = &cPtr
	incPtr.Increment()
	fmt.Println("Pointer Receiver:", cPtr.Value)
}
