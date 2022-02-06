package product

import (
	"database/sql"
	"fmt"
	"log"

	"sirclo/layered/relation/entities"
)

//jangan lupa inisialisasi new , berlaku di controller

type ProductRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProducts() ([]entities.Product, error) {
	var Products []entities.Product
	result, err := r.db.Query("select products.id, products.userid, products.nama, products.harga from products ")
	if err != nil {
		return Products, fmt.Errorf("error di dbquery")
	}
	defer result.Close()
	for result.Next() {
		var Product entities.Product
		err := result.Scan(&Product.Id,&Product.UserId.Id,&Product.Nama,&Product.Harga)
		if err != nil {
			log.Fatal("error di scan getProduct",err)
		}
		Products = append(Products, Product)
	}
	return Products, nil

}

func (r *ProductRepository)GetProduct(id int) (entities.Product, error) {
	var Product entities.Product
	result, err := r.db.Query("select id, userid, nama, harga from Products where id = ?", id)
	if err != nil {
		return Product, err
	}

	defer result.Close()
	for result.Next() {
		err := result.Scan(&Product.Id,&Product.UserId.Id,&Product.Nama,&Product.Harga)
		if err != nil {
			return Product, err
		}
		return Product, nil
	}
	return Product, fmt.Errorf("Product not found")
}

func (r *ProductRepository)UpdateProduct(id int, Product entities.Product) error {
	result, err := r.db.Exec("UPDATE Products SET nama= ?, harga= ? WHERE id = ?" , Product.Nama, Product.Harga, id)
	if err != nil {
		return err
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return fmt.Errorf("Product not found")
	}
	return nil
}

func (r *ProductRepository)DeleteProduct(id int) error {
	result, err := r.db.Exec("DELETE from Products where id = ?", id)
	if err != nil {
		return err
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return fmt.Errorf("Product not found")
	}
	return nil
}

func (r *ProductRepository)PostProduct(Product entities.Product) error {
	result, err := r.db.Exec("INSERT INTO Products(userid,nama ,harga) VALUES(?,?,?)", Product.UserId.Id, Product.Nama, Product.Harga)
	if err != nil {
		return err
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return fmt.Errorf("Product not created")
	}
	return nil
}