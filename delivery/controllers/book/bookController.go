package book

import (
	"net/http"
	"sirclo/layered/relation/delivery/common"
	"sirclo/layered/relation/delivery/middlewares"
	"sirclo/layered/relation/entities"
	"sirclo/layered/relation/repository/book"
	"strconv"

	"github.com/labstack/echo/v4"
)

// nah panggil interface di dalem controller
type BookController struct {
	repository book.BookInterface
}

func New(Book book.BookInterface) *BookController {
	return &BookController{
		repository : Book,
	}
}

func (uc BookController) GetBooksController() echo.HandlerFunc {
	return func(c echo.Context) error {
		Books,err := uc.repository.GetBooks()
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		return c.JSON(http.StatusOK, Books)
	}
}

// insert Book to DB
func (uc BookController) AddBookController() echo.HandlerFunc {
	return func(c echo.Context) error {
		if _, err := middlewares.ExtractTokenUserId(c); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		Book := entities.Book{}
		if err_bind := c.Bind(&Book); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}

		err := uc.repository.PostBook(Book)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		return c.JSON(http.StatusOK, common.SuccessOperation())
	}
}

func (uc BookController)GetBookController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// getting the id
		Bookid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// taking the data from DB
		var Book entities.Book
		Book, errsearch := uc.repository.GetBook(Bookid)
		if errsearch != nil {
			return c.JSON(http.StatusBadRequest, common.NotFound())
		}
		return c.JSON(http.StatusOK, common.SuccesOperationWithData(Book))
	}
}

// delete Book by id
func (uc BookController)DeleteBookController() echo.HandlerFunc {
	return func(c echo.Context) error {
		if _, err := middlewares.ExtractTokenUserId(c); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		Bookid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		if err := uc.repository.DeleteBook(Bookid); err != nil {
			return c.JSON(http.StatusBadRequest, common.NotFound())
		}

		return c.JSON(http.StatusOK, common.SuccessOperation())
	}
}

// update Book by id
func (uc BookController)UpdateBookController() echo.HandlerFunc {
	return func(c echo.Context) error {
		if _, err := middlewares.ExtractTokenUserId(c); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		Book := entities.Book{}
		if err_bind := c.Bind(&Book); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}
		// getting the id
		Bookid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		err_update := uc.repository.UpdateBook(Bookid, Book)
		if err_update != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation())
	}
}
