package book

import (
	"sirclo/layered/relation/entities"
)

type BookInterface interface{
	GetBooks() ([]entities.Book, error)
	PostBook(Book entities.Book) error 
	GetBook(id int) (entities.Book, error) 
	DeleteBook(id int) error
	UpdateBook(id int, Book entities.Book) error
}