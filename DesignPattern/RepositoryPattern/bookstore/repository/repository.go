package repository

import (
	"bookstore-repository/bookstore/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Repository interface (contract)
type BookRepository interface {
	GetByID(id int) (*models.Book, error)
	Save(book *models.Book) error
	GetAll() ([]*models.Book, error)
	Delete(id int) error
}

// Factory to switch DB
func GetRepository(dbType string, dsn string) (BookRepository, error) {
	switch dbType {
	case "memory":
		return NewInMemoryBookRepository(), nil
	case "mysql":
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS books (
			id INT PRIMARY KEY,
			title VARCHAR(255),
			author VARCHAR(255)
		)`)
		if err != nil {
			return nil, err
		}
		return NewMySQLBookRepository(db), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
