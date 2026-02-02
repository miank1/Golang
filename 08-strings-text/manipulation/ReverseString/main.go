package main

import "fmt"

func reverseString(s string) string {

	s1 := []byte(s)

	for i := 0; i < len(s1)/2; i++ {
		// Swap s1[i] and s1[len(s1)-1-i]
		s1[i], s1[len(s1)-1-i] = s1[len(s1)-1-i], s1[i]
	}
	return string(s1)
}

func main() {
	fmt.Println(reverseString("hello"))
}
