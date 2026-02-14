package main

import "fmt"

func areAnagrams(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	} else {

		m := make(map[rune]int)

		// Iterate over the first string and count occurrences of each rune
		for _, ch := range s1 {
			m[ch]++
		}

		// Iterate over the second string and decrease the count
		for _, ch := range s2 {
			m[ch]--
		}

		for _, count := range m {
			if count != 0 {
				return false
			}
		}
	}

	return true

}

func main() {
	fmt.Println(areAnagrams("listen", "silent"))
	fmt.Println(areAnagrams("hello", "world"))
}
