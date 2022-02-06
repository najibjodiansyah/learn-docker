package book

import (
	"database/sql"
	"fmt"
	"log"

	"sirclo/layered/relation/entities"
)

//jangan lupa inisialisasi new , berlaku di controller

type BookRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetBooks() ([]entities.Book, error) {
	var Books []entities.Book
	result, err := r.db.Query("select id, title, description, publisher from Books")
	if err != nil {
		return Books, err
	}
	defer result.Close()
	for result.Next() {
		var Book entities.Book
		err := result.Scan(&Book.Id,&Book.Title,&Book.Description,&Book.Publisher)
		if err != nil {
			log.Fatal("error di scan getBook")
		}
		Books = append(Books, Book)
	}
	return Books, nil

}

func (r *BookRepository)GetBook(id int) (entities.Book, error) {
	var Book entities.Book
	result, err := r.db.Query("select id, title, description, publisher from Books where id = ?", id)
	if err != nil {
		return Book, err
	}

	defer result.Close()
	for result.Next() {
		err := result.Scan(&Book.Id,&Book.Title,&Book.Description,&Book.Publisher)
		if err != nil {
			return Book, err
		}
		return Book, nil
	}
	return Book, fmt.Errorf("Book not found")
}

func (r *BookRepository)UpdateBook(id int, Book entities.Book) error {
	result, err := r.db.Exec("UPDATE Books SET title= ?, description= ?, publisher= ? WHERE id = ?" ,Book.Title, Book.Description, Book.Publisher, id)
	if err != nil {
		return err
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return fmt.Errorf("Book not found")
	}
	return nil
}

func (r *BookRepository)DeleteBook(id int) error {
	result, err := r.db.Exec("DELETE from Books where id = ?", id)
	if err != nil {
		return err
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return fmt.Errorf("Book not found")
	}
	return nil
}

func (r *BookRepository)PostBook(Book entities.Book) error {
	result, err := r.db.Exec("INSERT INTO Books(title, description, publisher) VALUES(?,?,?)", Book.Title, Book.Description, Book.Publisher)
	if err != nil {
		return err
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return fmt.Errorf("Book not created")
	}
	return nil
}