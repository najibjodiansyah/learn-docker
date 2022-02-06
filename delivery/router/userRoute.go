package router

import (
	"database/sql"
	"sirclo/layered/relation/delivery/controllers/user"

	"sirclo/layered/relation/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func InitUserRoute(e *echo.Echo, 
	userController *user.UserController,
	db *sql.DB) {
	e.GET("/users", userController.GetUsersController(),middlewares.JWTMiddleware())
	e.POST("/users", userController.AddUserController())
	e.GET("/users/:id", userController.GetUserController(),middlewares.JWTMiddleware())
	e.PUT("/users/:id", userController.UpdateUserController(),middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userController.DeleteUserController(),middlewares.JWTMiddleware())
}