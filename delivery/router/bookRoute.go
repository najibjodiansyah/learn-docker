package router

import (
	"database/sql"

	"sirclo/layered/relation/delivery/controllers/book"
	"sirclo/layered/relation/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func InitBookRoute(e *echo.Echo, 
	bookController *book.BookController,
	db *sql.DB) {
	
	e.GET("/books", bookController.GetBooksController())
	e.POST("/books", bookController.AddBookController(), middlewares.JWTMiddleware())
	e.GET("/books/:id", bookController.GetBookController())
	e.PUT("/books/:id", bookController.UpdateBookController(), middlewares.JWTMiddleware())
	e.DELETE("/books/:id", bookController.DeleteBookController(), middlewares.JWTMiddleware())
}