package main

import "fmt"

func main() {
	str := "Hello"
	fmt.Printf("Input string is %s \n", str)

	str = reverseString(str)
	fmt.Printf("Input string after reverse %s ", str)

}

func reverseString(str string) string {
	inputRune := []rune(str)

	l := 0
	r := len(inputRune) - 1

	for l < r {
		inputRune[l], inputRune[r] = inputRune[r], inputRune[l]
		l++
		r--
	}

	return string(inputRune)
}
