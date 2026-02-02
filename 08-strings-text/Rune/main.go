// Write a program in Golang that takes a string as input and counts the number of occurrences of each unique rune (character) in the string.
// Display the count of each rune along with the rune itself.

package main

import "fmt"

func countRunes(input string) map[rune]int {

	cnt := make(map[rune]int)

	for _, char := range input {
		cnt[char]++
	}

	return cnt
}

func main() {
	inputString := "Hello World"

	totalRune := countRunes(inputString)

	for x, v := range totalRune {
		fmt.Println(string(x), "-->", v)
	}

}
