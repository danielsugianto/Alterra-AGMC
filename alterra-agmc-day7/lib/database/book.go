package database

import (
	"context"
	"log"
	"strconv"

	"github.com/danielsugianto/alterra-agmc-day7/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDBBooksRepository struct {
	mongoDB *mongo.Database
}

// NewMongoDBBooksRepository will create an object that represent the users.Repository interface
func NewMongoDBBooksRepository(mongoDB *mongo.Database) models.BooksMongoDBRepository {
	return &mongoDBBooksRepository{mongoDB}
}

// get all books
func (mongoDBBooksRepository *mongoDBBooksRepository) GetBooks() ([]models.Book, error) {
	var books []models.Book
	findOptions := options.Find()

	csr, err := mongoDBBooksRepository.mongoDB.Collection("books").Find(context.Background(), findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(context.Background())

	for csr.Next(context.Background()) {
		var row models.Book
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		books = append(books, row)
	}
	return books, nil
}

// create new book
func (mongoDBBooksRepository *mongoDBBooksRepository) CreateBook(book models.Book) (models.Book, error) {
	_, err := mongoDBBooksRepository.mongoDB.Collection("books").InsertOne(context.Background(), book)
	if err != nil {
		return book, err
	}
	return book, nil
}

// get book by ID
func (mongoDBBooksRepository *mongoDBBooksRepository) GetBook(id string) (models.Book, error) {
	idInt, _ := strconv.Atoi(id)
	var getBook models.Book

	var selector = bson.M{"id": idInt}
	err := mongoDBBooksRepository.mongoDB.Collection("books").FindOne(context.Background(), selector).Decode(&getBook)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return getBook, nil
		}
		log.Fatal(err.Error())
	}
	return getBook, nil
}

// delete book by ID
func (mongoDBBooksRepository *mongoDBBooksRepository) DeleteBook(id string) error {
	idInt, _ := strconv.Atoi(id)
	var selector = bson.M{"id": idInt}
	_, err := mongoDBBooksRepository.mongoDB.Collection("books").DeleteOne(context.Background(), selector)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

// update book by ID
func (mongoDBBooksRepository *mongoDBBooksRepository) UpdateBook(id string, bookParam models.Book) (models.Book, error) {
	var book models.Book
	idInt, _ := strconv.Atoi(id)
	var selector = bson.M{"id": idInt}
	book.ID = idInt
	book.Name = bookParam.Name
	book.Year = bookParam.Year
	_, err := mongoDBBooksRepository.mongoDB.Collection("books").UpdateOne(context.Background(), selector, bson.M{"$set": book})
	if err != nil {
		log.Fatal(err.Error())
	}
	return book, nil
}
