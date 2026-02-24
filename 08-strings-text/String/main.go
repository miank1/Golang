package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	fmt.Printf("Length is %d \n", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	for i, letters := range str {
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letters))
	}

	fmt.Printf("Printing the character of string %c", str[0])

	// Repeat string
	copy := strings.Repeat("a", 2)
	fmt.Println(copy)

	copy = strings.Repeat("abc", 3)
	fmt.Println(copy)
}
