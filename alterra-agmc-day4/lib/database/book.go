package database

import (
	"strconv"

	"github.com/danielsugianto/alterra-agmc-day4/models"
)

var books []models.Book

func init() {
	books = append(books, models.Book{ID: 1, Name: "book 1", Year: 2010})
	books = append(books, models.Book{ID: 2, Name: "book 2", Year: 2011})
	books = append(books, models.Book{ID: 3, Name: "book 3", Year: 2012})
	books = append(books, models.Book{ID: 4, Name: "book 4", Year: 2013})
	books = append(books, models.Book{ID: 5, Name: "book 5", Year: 2014})
}

// get all books
func GetBooks() ([]models.Book, error) {
	return books, nil
}

// create new book
func CreateBook(book models.Book) (models.Book, error) {
	books = append(books, book)
	return book, nil
}

// get book by ID
func GetBook(id string) (models.Book, error) {
	idInt, _ := strconv.Atoi(id)
	var getBook models.Book

	for _, book := range books {
		if book.ID == idInt {
			getBook = book
			break
		}
	}
	return getBook, nil
}

// delete book by ID
func DeleteBook(id string) error {
	idInt, _ := strconv.Atoi(id)

	for i, book := range books {
		if book.ID == idInt {
			books[i] = books[len(books)-1]
			books[len(books)-1] = models.Book{}
			books = books[:len(books)-1]
			break
		}
	}
	return nil
}

// update book by ID
func UpdateBook(id string, bookParam models.Book) (models.Book, error) {
	var book models.Book
	idInt, _ := strconv.Atoi(id)
	for i := range books {
		if books[i].ID == idInt {
			books[i].Name = bookParam.Name
			books[i].Year = bookParam.Year
			book = books[i]
			break
		}
	}
	return book, nil
}
