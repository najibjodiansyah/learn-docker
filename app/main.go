package main

import (
	"fmt"
	"sirclo/layered/relation/config"
	"sirclo/layered/relation/delivery/middlewares"
	"sirclo/layered/relation/delivery/router"
	singleton "sirclo/layered/relation/util"

	_authController "sirclo/layered/relation/delivery/controllers/auth"
	_authRepo "sirclo/layered/relation/repository/auth"

	_bookController "sirclo/layered/relation/delivery/controllers/book"
	_bookRepo "sirclo/layered/relation/repository/book"

	_userController "sirclo/layered/relation/delivery/controllers/user"
	_userRepo "sirclo/layered/relation/repository/user"

	_productController "sirclo/layered/relation/delivery/controllers/product"
	_productRepo "sirclo/layered/relation/repository/product"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	
	config := config.GetConfig()

	db := singleton.MysqlDriver(config)	

	//initiate user model
	authRepo := _authRepo.New(db)
	bookRepo := _bookRepo.New(db)
	userRepo := _userRepo.New(db)
	productRepo := _productRepo.New(db)

	//initiate user controller
	authController := _authController.New(authRepo)
	bookController := _bookController.New(bookRepo)
	userController := _userController.New(userRepo)
	productController := _productController.New(productRepo)

	e := echo.New()

	middlewares.LogMiddlewares(e)

	router.InitLoginRoute(e,authController,db)
	router.InitBookRoute(e,bookController,db)
	router.InitUserRoute(e,userController,db)
	router.InitProductRoute(e,productController,db)
	
	address := fmt.Sprintf("localhost:%d", config.Port)
	e.Logger.Fatal(e.Start(address))
}

/* 1 buat entities
	2 mau ngapain aja -> init = repository -> buat interface 
	3 hasil dari interface di panggil di controller -> masuk ke route 
	4 setting di main pemanggilan route
*/