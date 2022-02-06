package router

import (
	"database/sql"
	"sirclo/layered/relation/delivery/controllers/auth"

	"github.com/labstack/echo/v4"
)

func InitLoginRoute(e *echo.Echo, 
	authController *auth.AuthController,
	db *sql.DB) {

	e.POST("/login", authController.CreateLogin())
}
