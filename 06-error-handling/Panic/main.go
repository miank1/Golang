package main

import "fmt"

func countFrequency(input string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if input == "" {
		panic("input string cannot be empty")
	}

	freqMap := make(map[rune]int)

	for _, ch := range input {
		freqMap[ch]++
	}

	for k, v := range freqMap {
		fmt.Printf("%q: %d\n", k, v)
	}
}
func main() {
	countFrequency("golang")

	fmt.Println("-----------------------")
	countFrequency("") // this will panic, but recover will catch it.
}
