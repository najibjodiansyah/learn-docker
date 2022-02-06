package product

import (
	"fmt"
	"net/http"
	"sirclo/layered/relation/delivery/common"
	"sirclo/layered/relation/delivery/middlewares"
	"sirclo/layered/relation/entities"
	"sirclo/layered/relation/repository/product"
	"strconv"

	"github.com/labstack/echo/v4"
)

// nah panggil interface di dalem controller
type ProductController struct {
	repository product.ProductInterface
}

func New(Product product.ProductInterface) *ProductController {
	return &ProductController{
		repository : Product,
	}
}

func (uc ProductController) GetProductsController() echo.HandlerFunc {
	return func(c echo.Context) error {
		Products,err := uc.repository.GetProducts()
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		return c.JSON(http.StatusOK, Products)
	}
}

// CREATE ADD product
func (uc ProductController) AddProductController() echo.HandlerFunc {
	return func(c echo.Context) error {
		//cek current user dari jwt
		currentUser, err := middlewares.ExtractTokenUserId(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		//cek bind produk
		Product := entities.Product{}
		err_bind := c.Bind(&Product)
		if  err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}
		Product.UserId.Id = currentUser 
		fmt.Println(currentUser, Product.UserId.Id, err_bind)
		//run sql sintaks
		if err := uc.repository.PostProduct(Product); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation())
	}
}

// get product by id
func (uc ProductController)GetProductController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// getting the id
		Productid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// taking the data from DB
		var Product entities.Product
		Product, errsearch := uc.repository.GetProduct(Productid)
		if errsearch != nil {
			return c.JSON(http.StatusBadRequest, common.NotFound())
		}
		return c.JSON(http.StatusOK, common.SuccesOperationWithData(Product))
	}
}

// delete Product by id
func (uc ProductController)DeleteProductController() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser, err := middlewares.ExtractTokenUserId(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		Productid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// var Product entities.Product
		Product, _ := uc.repository.GetProduct(Productid)
		fmt.Println(Product)
		if Product.UserId.Id != currentUser {
			return c.JSON(http.StatusBadRequest, common.NotFound())
		}
		if err := uc.repository.DeleteProduct(Productid); err != nil {
			return c.JSON(http.StatusBadRequest, common.NotFound())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation())
	}
}

// update Product by id
func (uc ProductController)UpdateProductController() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser, err := middlewares.ExtractTokenUserId(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		Product := entities.Product{}
		if err_bind := c.Bind(&Product); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}
		// getting the id
		Productid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		ProductUserId, _ := uc.repository.GetProduct(Productid)
		fmt.Println(Product)
		if ProductUserId.UserId.Id != currentUser {
			return c.JSON(http.StatusBadRequest, common.NotFound())
		}
		err_update := uc.repository.UpdateProduct(Productid, Product)
		if err_update != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation())
	}
}
