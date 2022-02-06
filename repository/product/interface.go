package product

import (
	"sirclo/layered/relation/entities"
)

type ProductInterface interface{
	GetProducts() ([]entities.Product, error)
	PostProduct(Product entities.Product) error 
	GetProduct(id int) (entities.Product, error) 
	DeleteProduct(id int) error
	UpdateProduct(id int, Product entities.Product) error
}