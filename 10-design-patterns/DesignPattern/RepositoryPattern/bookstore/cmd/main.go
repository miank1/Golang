package main

import (
	"bookstore-repository/bookstore/config"
	"bookstore-repository/bookstore/models"
	"bookstore-repository/bookstore/repository"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()

	repo, err := repository.GetRepository(cfg.DBType, cfg.DSN)
	if err != nil {
		panic(err)
	}

	// Save books
	repo.Save(&models.Book{ID: 1, Title: "Harry Potter", Author: "J.K. Rowling"})
	repo.Save(&models.Book{ID: 2, Title: "The Hobbit", Author: "J.R.R. Tolkien"})

	// Fetch one
	book, _ := repo.GetByID(1)
	fmt.Println("Found Book:", book.Title, "by", book.Author)

	// Fetch all
	allBooks, _ := repo.GetAll()
	fmt.Println("All Books in", cfg.DBType, "DB:")
	for _, b := range allBooks {
		fmt.Println("-", b.Title)
	}
}
