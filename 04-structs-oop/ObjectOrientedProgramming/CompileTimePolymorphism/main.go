package main

import "fmt"

// The compile time polymorphism is achieved using vardiac functions in golang

type maths struct {
}

func (m *maths) add(numbers ...int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func main() {
	m := &maths{}
	fmt.Printf("Result := %d\n", m.add(1, 2))
	fmt.Printf("Result := %d\n", m.add(1, 2, 3))

}
