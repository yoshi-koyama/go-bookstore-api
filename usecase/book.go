package usecase

import (
	"context"
	"bookstore-api/domain/repository"
	"errors"
)

type Book interface {
	BuyBooks(ctx context.Context, id int, amountToBuy int) (*string, error)
	GetAllBooks(ctx context.Context) ([]BookDTO, error)
	GetBook(ctx context.Context, id int) (*BookDTO, error)
}

type BookDTO struct {
	ID    int
	Name  string
	Price int
}

type bookUseCase struct {
	bookRepo    repository.Book
	paymentRepo repository.Payment
}

func NewBook(bookRepo repository.Book, paymentRepo repository.Payment) Book {
	return bookUseCase{
		bookRepo:    bookRepo,
		paymentRepo: paymentRepo,
	}
}

func (b bookUseCase) BuyBooks(ctx context.Context, id int, amountToBuy int) (*string, error) {
	book := b.bookRepo.FindByID(id)
	if book == nil {
		return nil, errors.New("cannot find book")
	}

	result := b.paymentRepo.MakePayment(book.Price() * amountToBuy)
	return &result, nil
}

func (b bookUseCase) GetAllBooks(ctx context.Context) ([]BookDTO, error) {
	books := b.bookRepo.FindAll()

	var result []BookDTO
	for _, book := range books {
		result = append(result, BookDTO{
			ID:    book.ID(),
			Name:  book.Name(),
			Price: book.Price(),
		})
	}

	return result, nil
}

func (b bookUseCase) GetBook(ctx context.Context, id int) (*BookDTO, error) {
	book := b.bookRepo.FindByID(id)
	if book == nil {
		return nil, errors.New("cannot find book")
	}

	result := BookDTO{
		ID:    book.ID(),
		Name:  book.Name(),
		Price: book.Price(),
	}

	return &result, nil
}
