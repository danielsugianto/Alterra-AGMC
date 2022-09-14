package controllers

import (
	"net/http"

	"github.com/danielsugianto/alterra-agmc-day3/lib/database"
	"github.com/danielsugianto/alterra-agmc-day3/models"
	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book
	books, err := database.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	if err := c.Validate(book); err != nil {
		return err
	}
	newBook, err := database.CreateBook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    newBook,
	})
}

// get book by ID
func GetBookController(c echo.Context) error {
	id := c.Param("id")
	var book models.Book

	book, err := database.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

// delete book by ID
func DeleteBookController(c echo.Context) error {
	id := c.Param("id")

	err := database.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update book by ID
func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	bookParam := models.Book{}
	c.Bind(&bookParam)
	var book models.Book

	book, err := database.UpdateBook(id, bookParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}
