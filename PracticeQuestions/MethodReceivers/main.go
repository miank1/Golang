package main

import "fmt"

type Counter struct {
	Value int
}

func (c Counter) IncrementByValue() {
	c.Value += 1
}

func (c *Counter) IncrementByPointer() {
	c.Value += 1
}

func main() {

	c := Counter{
		Value: 0,
	}
	c.IncrementByValue()
	fmt.Println(c.Value)
	c.IncrementByPointer()
	fmt.Println(c.Value)
}
