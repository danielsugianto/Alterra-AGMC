package models

import "github.com/labstack/echo/v4"

type Book struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" validate:"required" bson:"name"`
	Year int    `json:"year" validate:"required" bson:"year"`
}

type BooksUsecase interface {
	GetBooksUsecase() ([]Book, error)
	CreateBookUsecase(bookParam Book) (Book, error)
	GetBookUsecase(id string) (Book, error)
	DeleteBookUsecase(id string) error
	UpdateBookUsecase(id string, bookParam Book) (Book, error)
}

type BooksMongoDBRepository interface {
	GetBooks() ([]Book, error)
	CreateBook(userParam Book) (Book, error)
	GetBook(id string) (Book, error)
	DeleteBook(id string) error
	UpdateBook(id string, userParam Book) (Book, error)
}

type BooksController interface {
	GetBooksController(c echo.Context) error
	CreateBookController(c echo.Context) error
	GetBookController(c echo.Context) error
	DeleteBookController(c echo.Context) error
	UpdateBookController(c echo.Context) error
}
