package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"phase3-gc1-shopping/helper"
	"phase3-gc1-shopping/model/web"
	"phase3-gc1-shopping/service"
)

type ProductController interface {
	CreateProduct(c echo.Context) error
	GetProductById(c echo.Context) error
	GetProducts(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

type ProductControllerImpl struct {
	service.ProductService
}

func NewProductController(productService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{ProductService: productService}
}

func (controller *ProductControllerImpl) CreateProduct(c echo.Context) error {
	var request web.ProductRequest
	err := c.Bind(&request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		return c.JSON(errResponse.Code, errResponse)
	}
	response, errResponse := controller.ProductService.CreateProduct(request)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusCreated, web.Response{
		Message: "success create product",
		Data:    response,
	})
}

func (controller *ProductControllerImpl) GetProductById(c echo.Context) error {
	productID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errResponse := helper.ErrBadRequest(err)
		return c.JSON(errResponse.Code, errResponse)
	}
	response, errResponse := controller.ProductService.GetProductById(productID)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusOK, web.Response{
		Message: "success get product",
		Data:    response,
	})
}

func (controller *ProductControllerImpl) GetProducts(c echo.Context) error {
	responses, errResponse := controller.ProductService.GetProducts()
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusOK, web.Response{
		Message: "success get all products",
		Data:    responses,
	})
}

func (controller *ProductControllerImpl) UpdateProduct(c echo.Context) error {
	productID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errResponse := helper.ErrBadRequest(err)
		return c.JSON(errResponse.Code, errResponse)
	}
	var request web.ProductRequest
	err = c.Bind(&request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		return c.JSON(errResponse.Code, errResponse)
	}
	transaction, errResponse := controller.ProductService.UpdateProduct(request, productID)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusOK, web.Response{
		Message: "success update product",
		Data:    transaction,
	})
}

func (controller *ProductControllerImpl) DeleteProduct(c echo.Context) error {
	productID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errResponse := helper.ErrBadRequest(err)
		return c.JSON(errResponse.Code, errResponse)
	}
	errResponse := controller.ProductService.DeleteProduct(productID)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusOK, web.Response{
		Message: "success delete product",
	})
}
