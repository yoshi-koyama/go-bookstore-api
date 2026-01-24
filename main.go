package main

import (
	"dependency-injection-sample/handler"
	"dependency-injection-sample/infra/dao"
	"dependency-injection-sample/infra/external"
	"dependency-injection-sample/usecase"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	bookRepo := dao.NewBook()
	rakutenPay := external.NewRakutenPay()
	bookUseCase := usecase.NewBook(bookRepo, rakutenPay)
	bookHandler := handler.NewBookHandler(bookUseCase)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Route("/bookstore/api", func(r chi.Router) {
		r.Get("/books", bookHandler.GetBooks)
		r.Post("/checkouts", bookHandler.Checkout)
	})

	http.ListenAndServe(":8080", r)
}
