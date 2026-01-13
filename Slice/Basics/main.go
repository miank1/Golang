// Slice basics

package main

import "fmt"

func main() {
	s := make([]int, 0, 4)

	fmt.Println("Length ", len(s))
	fmt.Println("Capacity ", cap(s))

	// Appending Slice

	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)

	// Capacity exceeded - new array allocated and old data copied - slice now points to new array.

	s = append(s, 5)
	s = append(s, 6)

	fmt.Println("Length ", len(s))
	fmt.Println("Capacity ", cap(s))

}
