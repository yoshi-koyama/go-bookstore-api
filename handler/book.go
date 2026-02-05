package handler

import (
	"bookstore-api/handler/request"
	"bookstore-api/handler/response"
	"bookstore-api/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type Book interface {
	Checkout(w http.ResponseWriter, r *http.Request)
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
}

type bookHandler struct {
	useCase usecase.Book
}

func NewBookHandler(useCase usecase.Book) Book {
	return bookHandler{
		useCase: useCase,
	}
}

func (b bookHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.BuyBooks
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	result, err := b.useCase.BuyBooks(ctx, *req.ID, *req.AmountToPay)

	if err != nil {
		if err.Error() == "cannot find book" {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{
				"message": "no book found",
			})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"message": "something went wrong",
		})
		return

	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, response.NewBuyBooks(*result))
}

func (b bookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	books, err := b.useCase.GetAllBooks(ctx)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"message": "something went wrong",
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, response.NewBookList(books))
}

func (b bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// パスパラメータからIDを取得
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"message": "invalid book id",
		})
		return
	}

	// UseCaseを呼び出し
	book, err := b.useCase.GetBook(ctx, id)
	if err != nil {
		if err.Error() == "cannot find book" {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{
				"message": "no book found",
			})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"message": "something went wrong",
		})
		return
	}

	// レスポンスを返す
	render.Status(r, http.StatusOK)
	render.Render(w, r, response.NewBook(*book))
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
