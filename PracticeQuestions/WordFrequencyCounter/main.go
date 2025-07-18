package main

import (
	"fmt"
	"strings"
)

func WordFrequency(str string) map[string]int {

	words := strings.Split(str, " ")
	m1 := make(map[string]int)

	for _, word := range words {
		m1[word]++
	}

	return m1
}

func main() {
	text := "go is fun and go is powerful"
	freq := WordFrequency(text)
	fmt.Println(freq)
}
