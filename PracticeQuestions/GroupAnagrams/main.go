package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str) // sort the characters
		fmt.Println("------------------", sortedStr)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	fmt.Println("Map", anagramMap)

	// Extract grouped anagrams
	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

// Helper to sort characters of a string
func sortString(s string) string {
	chars := strings.Split(s, "")
	fmt.Println(chars)
	sort.Strings(chars)
	//fmt.Println("---------", chars)

	return strings.Join(chars, "")
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(input)

	for _, group := range result {
		fmt.Println(group)
	}
}
