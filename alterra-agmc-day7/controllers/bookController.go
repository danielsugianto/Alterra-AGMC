package controllers

import (
	"net/http"
	"strconv"

	"github.com/danielsugianto/alterra-agmc-day7/models"
	"github.com/labstack/echo/v4"
)

type booksController struct {
	booksUsecase models.BooksUsecase
}

// NewBooksUsecase will create new an booksController object representation of domain.BooksUsecase interface
func NewBooksController(booksUsecase models.BooksUsecase) models.BooksController {
	return &booksController{
		booksUsecase: booksUsecase,
	}
}

// get all books
func (booksController *booksController) GetBooksController(c echo.Context) error {
	var books []models.Book
	books, err := booksController.booksUsecase.GetBooksUsecase()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// create new book
func (booksController *booksController) CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	if err := c.Validate(book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newBook, err := booksController.booksUsecase.CreateBookUsecase(book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new book",
		"book":    newBook,
	})
}

// get book by ID
func (booksController *booksController) GetBookController(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	_, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "Can't Get Book")
	}
	book, err := booksController.booksUsecase.GetBookUsecase(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// delete book by ID
func (booksController *booksController) DeleteBookController(c echo.Context) error {
	id := c.Param("id")

	_, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "Can't Delete Book")
	}
	err := booksController.booksUsecase.DeleteBookUsecase(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update book by ID
func (booksController *booksController) UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	bookParam := models.Book{}
	c.Bind(&bookParam)
	if err := c.Validate(bookParam); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var book models.Book

	book, err := booksController.booksUsecase.UpdateBookUsecase(id, bookParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}
