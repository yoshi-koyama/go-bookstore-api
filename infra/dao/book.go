package dao

import (
	"bookstore-api/domain/model"
	"bookstore-api/domain/repository"

	"github.com/jmoiron/sqlx"
)

type bookRepository struct {
	db *sqlx.DB
}

func NewBook(db *sqlx.DB) repository.Book {
	return bookRepository{
		db: db,
	}
}

func (b bookRepository) FindByID(id int) *model.Book {
	var record bookRecord
	err := b.db.Get(&record, "SELECT id, name, price FROM books WHERE id = ?", id)
	if err != nil {
		return nil
	}

	book := model.NewBook(record.ID, record.Name, record.Price)
	return &book
}

type bookRecord struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Price int    `db:"price"`
}

func (b bookRepository) FindAll() []model.Book {
	var records []bookRecord
	err := b.db.Select(&records, "SELECT id, name, price FROM books")
	if err != nil {
		return []model.Book{}
	}

	books := make([]model.Book, 0, len(records))
	for _, record := range records {
		books = append(books, model.NewBook(record.ID, record.Name, record.Price))
	}

	return books
}
