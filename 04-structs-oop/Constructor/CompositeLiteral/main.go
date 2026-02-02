package main

import "fmt"

// using composite literal

type Movie struct {
	title  string
	rating int
}

func (m *Movie) printTitle() {
	fmt.Printf("Movie title is : %s \n", m.title)
}

func (m *Movie) printRating() {
	fmt.Printf("Rating is : %d \n", m.rating)
}

func main() {
	m := Movie{
		title:  "Top Gun",
		rating: 5,
	}

	m.printTitle()
	m.printRating()
}
