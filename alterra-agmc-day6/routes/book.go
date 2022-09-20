package routes

import (
	"github.com/danielsugianto/alterra-agmc-day6/models"
	"github.com/labstack/echo/v4"
)

func UnauthenticatedBookRoutes(e *echo.Group, booksController models.BooksController) {
	e.GET("/books", booksController.GetBooksController)
	e.GET("/books/:id", booksController.GetBookController)
}

func AuthenticatedBookRoutes(e *echo.Group, booksController models.BooksController) {
	e.POST("/books", booksController.CreateBookController)
	e.DELETE("/books/:id", booksController.DeleteBookController)
	e.PUT("/books/:id", booksController.UpdateBookController)
}
