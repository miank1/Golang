package main

import "fmt"

type Movie struct {
	title  string
	rating int
}

// Creating using Init Method

func (m *Movie) init(title string, rating int) {
	m.title = title
	m.rating = rating
}

func (m *Movie) printTitle() {
	fmt.Printf("Movie title is :%s\n", m.title)
}

func (m *Movie) printRating() {
	fmt.Printf("Rating is :%d \n", m.rating)
}

func main() {
	m := new(Movie) // Movie{}
	m.init("Free Guy", 5)
	m.printTitle()
	m.printRating()
}
