package repository

import "bookstore-api/domain/model"

type Book interface {
	FindByID(id int) *model.Book
	FindAll() []model.Book
}
