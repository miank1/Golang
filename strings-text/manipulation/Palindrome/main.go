package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	str := []byte(s)
	i, j := 0, len(str)-1

	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	str := "madam"
	fmt.Println(isPalindrome(str)) // true
}
