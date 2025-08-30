package main

import (
	"fmt"
	"strings"
)

func wordFrequency(lines []string) {
	// Your code here

	for _, sentence := range lines {
		// Split sentence into words and append to the result
		lines = append(lines, strings.Fields(sentence)...)
	}

	fmt.Println(lines)

}

func main() {

	urls := []string{"The quick brown fox", "jumps over the lazy dog", "The FOX jumps again"}

	wordFrequency(urls)

}
