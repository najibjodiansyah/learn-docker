package router

import (
	"database/sql"

	"sirclo/layered/relation/delivery/controllers/product"
	"sirclo/layered/relation/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func InitProductRoute(e *echo.Echo, 
	productController *product.ProductController,
	db *sql.DB) {
	
	e.GET("/products", productController.GetProductsController())
	e.POST("/products", productController.AddProductController(), middlewares.JWTMiddleware())
	e.GET("/products/:id", productController.GetProductController())
	e.PUT("/products/:id", productController.UpdateProductController(), middlewares.JWTMiddleware())
	e.DELETE("/products/:id", productController.DeleteProductController(), middlewares.JWTMiddleware())
}