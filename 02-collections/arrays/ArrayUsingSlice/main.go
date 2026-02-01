package main

import (
	"fmt"
)

type Stack struct {
	data []interface{}
}

// NewStack initializes and returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

// Pop removes and returns the item from the top of the stack.
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]
	return item, nil
}

// Peek returns the item at the top of the stack without removing it.
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}

	return s.data[len(s.data)-1], nil
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack:", stack.data)

	top, err := stack.Peek()
	if err == nil {
		fmt.Println("Top of the stack:", top)
	}

	pop1, err := stack.Pop()
	if err == nil {
		fmt.Println("Popped item:", pop1)
	}

	fmt.Println("Stack after popping:", stack.data)
}
