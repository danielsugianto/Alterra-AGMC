package models

type Book struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Year int    `json:"year"`
}
