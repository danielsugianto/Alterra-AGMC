package models

type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
}
