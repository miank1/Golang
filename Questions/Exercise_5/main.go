package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func main() {
	c := Counter{
		Value: 0,
	}

	c.Increment()
	c.Increment()
	c.Increment()

	fmt.Println(c.Value)

}
