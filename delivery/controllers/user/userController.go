package user

import (
	"fmt"
	"net/http"
	"sirclo/layered/relation/delivery/common"
	"sirclo/layered/relation/delivery/middlewares"
	"sirclo/layered/relation/entities"
	"sirclo/layered/relation/repository/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

// nah panggil interface di dalem controller
type UserController struct {
	repository user.UserInterface
}

func New(user user.UserInterface) *UserController {
	return &UserController{
		repository : user,
	}
}

func (uc UserController) GetUsersController() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser, err := middlewares.ExtractTokenUserId(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		fmt.Println(currentUser)
		users,err := uc.repository.GetUsers()
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		return c.JSON(http.StatusOK, users)
	}
}

// insert user to DB
func (uc UserController) AddUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := entities.User{}
		if err_bind := c.Bind(&user); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}

		err := uc.repository.PostUser(user)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		return c.JSON(http.StatusOK, common.SuccessOperation())
	}
}

func (uc UserController)GetUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser, err := middlewares.ExtractTokenUserId(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		fmt.Println(currentUser)
		// getting the id
		userid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// taking the data from DB
		var user entities.User
		user, errsearch := uc.repository.GetUser(userid)
		if errsearch != nil {
			return c.JSON(http.StatusBadRequest, common.NotFound())
		}
		return c.JSON(http.StatusOK, common.SuccesOperationWithData(user))
	}
}

// delete user by id
func (uc UserController)DeleteUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser, err := middlewares.ExtractTokenUserId(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		fmt.Println(currentUser)

		userid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		if err := uc.repository.DeleteUser(userid); err != nil {
			return c.JSON(http.StatusBadRequest, common.NotFound())
		}

		return c.JSON(http.StatusOK, common.SuccessOperation())
	}
}

// update user by id
func (uc UserController)UpdateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser, err := middlewares.ExtractTokenUserId(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		fmt.Println(currentUser)

		user := entities.User{}
		if err_bind := c.Bind(&user); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}
		// getting the id
		userid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		err_update := uc.repository.UpdateUser(userid, user)
		if err_update != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation())
	}
}
