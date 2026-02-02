package repository

import (
	"bookstore-repository/bookstore/models"
	"database/sql"
	"fmt"
)

type MySQLBookRepository struct {
	db *sql.DB
}

func NewMySQLBookRepository(db *sql.DB) *MySQLBookRepository {
	return &MySQLBookRepository{db: db}
}

func (r *MySQLBookRepository) GetByID(id int) (*models.Book, error) {
	row := r.db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id)
	book := &models.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		return nil, fmt.Errorf("book not found: %w", err)
	}
	return book, nil
}

func (r *MySQLBookRepository) Save(book *models.Book) error {
	_, err := r.db.Exec("INSERT INTO books (id, title, author) VALUES (?, ?, ?)", book.ID, book.Title, book.Author)
	return err
}

func (r *MySQLBookRepository) GetAll() ([]*models.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		b := &models.Book{}
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *MySQLBookRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
