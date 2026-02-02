package main

import (
	"fmt"
	"strings"
)

func countVowels(s string) int {
	// Your code here

	s = strings.ToLower(s)

	count := 0

	r := []rune(s)

	for i := 0; i < len(r)-1; i++ {
		str := string(r[i])

		if str == "a" || str == "e" || str == "i" || str == "o" || str == "u" {
			count++
		}
	}

	return count
}

func main() {
	str := "Hello World !! Lets Code Golang !!"

	fmt.Println(countVowels(str))

}
