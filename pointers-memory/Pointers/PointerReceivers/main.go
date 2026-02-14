package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func (c Counter) GetValue() int {
	return c.Value
}

func main() {
	c := Counter{
		Value: 1,
	}

	c.Increment()
	c.Increment()

	fmt.Println("Value is", c.GetValue())
}
