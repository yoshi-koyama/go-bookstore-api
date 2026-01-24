package main

import (
	"bookstore-api/config"
	"bookstore-api/handler"
	"bookstore-api/infra/dao"
	"bookstore-api/infra/database"
	"bookstore-api/infra/external"
	"bookstore-api/usecase"
	"log"

	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	// 設定の読み込み
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// データベース接続の初期化
	db, err := database.NewDB(cfg)
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

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
