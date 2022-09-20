package usecase

import "github.com/danielsugianto/alterra-agmc-day6/models"

type booksUsecase struct {
	booksMongoDBRepo models.BooksMongoDBRepository
}

// NewBooksUsecase will create new an booksUsecase object representation of domain.BooksUsecase interface
func NewBooksUsecase(booksMongoDBRepo models.BooksMongoDBRepository) models.BooksUsecase {
	return &booksUsecase{
		booksMongoDBRepo: booksMongoDBRepo,
	}
}

// get all books
func (booksUsecase *booksUsecase) GetBooksUsecase() ([]models.Book, error) {
	var books []models.Book

	books, err := booksUsecase.booksMongoDBRepo.GetBooks()
	if err != nil {
		return books, err
	}
	return books, nil
}

// create new book
func (booksUsecase *booksUsecase) CreateBookUsecase(bookParam models.Book) (models.Book, error) {
	book, err := booksUsecase.booksMongoDBRepo.CreateBook(bookParam)
	if err != nil {
		return book, err
	}
	return book, nil
}

// get book by ID
func (booksUsecase *booksUsecase) GetBookUsecase(id string) (models.Book, error) {
	book, err := booksUsecase.booksMongoDBRepo.GetBook(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

// delete book by ID
func (booksUsecase *booksUsecase) DeleteBookUsecase(id string) error {
	errDelete := booksUsecase.booksMongoDBRepo.DeleteBook(id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

// update book by ID
func (booksUsecase *booksUsecase) UpdateBookUsecase(id string, bookParam models.Book) (models.Book, error) {
	book, errUpdate := booksUsecase.booksMongoDBRepo.UpdateBook(id, bookParam)
	if errUpdate != nil {
		return book, errUpdate
	}
	return book, nil
}
