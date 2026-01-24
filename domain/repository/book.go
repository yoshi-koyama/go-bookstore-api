package repository

import "dependency-injection-sample/domain/model"

type Book interface {
	FindByID(id int) *model.Book
	FindAll() []model.Book
}
