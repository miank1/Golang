package main

import "fmt"

func main() {
	words := []string{"i", "love", "golang", "i", "love", "coding"}
	k := 2

	fmt.Println(words, k)

	m1 := make(map[string]int)

	for i := 0; i < len(words); i++ {
		m1[words[i]]++
	}

	fmt.Println("map is ", m1)

}
