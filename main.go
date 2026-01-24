package main

import (
	"dependency-injection-sample/handler"
	"dependency-injection-sample/infra/dao"
	"dependency-injection-sample/infra/database"
	"dependency-injection-sample/infra/external"
	"dependency-injection-sample/usecase"
	"log"

	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	// データベース接続の初期化
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := chi.NewRouter()
	bookRepo := dao.NewBook(db)
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
