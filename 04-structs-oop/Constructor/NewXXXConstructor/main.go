package main

import "fmt"

type Movie struct {
	title  string
	rating int
}

func (m *Movie) printTitle() {
	fmt.Printf("Movie title is :%s \n", m.title)
}
func (m *Movie) printRating() {
	fmt.Printf("Rating is :%d \n", m.rating)
}

// Using NewXXX Constructor
func NewMovie(title string, rating int) *Movie {
	m := &Movie{
		title:  title,
		rating: rating,
	}
	return m
}

func main() {
	m := NewMovie("Top Gun Maverick", 5)
	m.printTitle()
	m.printRating()
}
