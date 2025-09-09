package repository

import (
	"bookstore-repository/bookstore/models"
	"fmt"
)

type InMemoryBookRepository struct {
	storage map[int]*models.Book
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		storage: make(map[int]*models.Book),
	}
}

func (r *InMemoryBookRepository) GetByID(id int) (*models.Book, error) {
	if book, ok := r.storage[id]; ok {
		return book, nil
	}
	return nil, fmt.Errorf("book not found")
}

func (r *InMemoryBookRepository) Save(book *models.Book) error {
	r.storage[book.ID] = book
	return nil
}

func (r *InMemoryBookRepository) GetAll() ([]*models.Book, error) {
	books := make([]*models.Book, 0, len(r.storage))
	for _, book := range r.storage {
		books = append(books, book)
	}
	return books, nil
}

func (r *InMemoryBookRepository) Delete(id int) error {
	if _, ok := r.storage[id]; !ok {
		return fmt.Errorf("book not found")
	}
	delete(r.storage, id)
	return nil
}
