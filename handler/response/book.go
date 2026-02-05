package response

import (
	"bookstore-api/usecase"
	"net/http"
)

type BuyBooks struct {
	Message string `json:"message"`
}

func NewBuyBooks(message string) *BuyBooks {
	return &BuyBooks{
		Message: message,
	}
}

func (b *BuyBooks) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Book struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func NewBook(book usecase.BookDTO) *Book {
	return &Book{
		ID:    book.ID,
		Name:  book.Name,
		Price: book.Price,
	}
}

func (b *Book) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type BookList struct {
	Books []Book `json:"books"`
}

func NewBookList(books []usecase.BookDTO) *BookList {
	var result []Book
	for _, book := range books {
		result = append(result, Book{
			ID:    book.ID,
			Name:  book.Name,
			Price: book.Price,
		})
	}
	return &BookList{
		Books: result,
	}
}

func (b *BookList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
