package main

import "fmt"

func FirstUniqueChar(s string) rune {

	if len(s) == 0 {
		return -1
	}

	freq := make(map[rune]int)

	for _, ch := range s {
		freq[ch]++
	}

	for _, ch := range s {
		if freq[ch] == 1 {
			return ch
		}
	}

	return -1
}

func main() {

	s := "swiss"

	s1 := FirstUniqueChar(s)

	fmt.Println(string(s1))

}
