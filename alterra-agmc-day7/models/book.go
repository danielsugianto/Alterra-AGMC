package models

type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
	Year int    `json:"year" validate:"required"`
}
