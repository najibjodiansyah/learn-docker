package auth

import (
	"fmt"
	"net/http"
	"sirclo/layered/relation/delivery/common"
	"sirclo/layered/relation/delivery/middlewares"
	"sirclo/layered/relation/repository/auth"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repository auth.AuthInterface
}

func New(Auth auth.AuthInterface) *AuthController {
	return &AuthController{
		repository : Auth,
	}
}

func (ac AuthController)CreateLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
	  // validasi user
	  var identity common.LoginRequestFormat
	  if err := c.Bind(&identity); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
		  "code":    400,
		  "status":  "failed",
		  "message": "failed to bind data",
		})
	  }
	  user, err :=  ac.repository.Login(identity)
	  if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		  "code":    401,
		  "status":  "unauthorized",
		  "message": "unauthorized access",
		})
	  }
	  // create token
	  token, err := middlewares.CreateToken(user.Id, user.Email)
	  if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		  "code":    500,
		  "status":  "internal server error",
		  "message": "cannot create token",
		})
	  }
  
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	  })
	}
  }