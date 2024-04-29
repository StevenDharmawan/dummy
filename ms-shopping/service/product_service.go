package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"phase3-gc1-shopping/helper"
	"phase3-gc1-shopping/model/domain"
	"phase3-gc1-shopping/model/web"
	"phase3-gc1-shopping/repository"
)

type ProductService interface {
	CreateProduct(request web.ProductRequest) (*domain.Product, *web.ErrorResponse)
	GetProducts() ([]domain.Product, *web.ErrorResponse)
	GetProductById(productID primitive.ObjectID) (*domain.Product, *web.ErrorResponse)
	UpdateProduct(request web.ProductRequest, productID primitive.ObjectID) (*domain.Product, *web.ErrorResponse)
	DeleteProduct(productID primitive.ObjectID) *web.ErrorResponse
}

type ProductServiceImpl struct {
	repository.ProductRepository
	*validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) CreateProduct(request web.ProductRequest) (*domain.Product, *web.ErrorResponse) {
	err := service.Validate.Struct(request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		fmt.Println(&errResponse)
		return nil, &errResponse
	}
	product := domain.Product{
		Name:  request.Name,
		Price: request.Price,
	}
	product, err = service.ProductRepository.CreateProduct(product)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return &product, nil
}

func (service *ProductServiceImpl) GetProducts() ([]domain.Product, *web.ErrorResponse) {
	products, err := service.ProductRepository.GetAllProducts()
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return products, nil
}

func (service *ProductServiceImpl) GetProductById(productID primitive.ObjectID) (*domain.Product, *web.ErrorResponse) {
	product, err := service.ProductRepository.GetProductById(productID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return &product, nil
}

func (service *ProductServiceImpl) UpdateProduct(request web.ProductRequest, productID primitive.ObjectID) (*domain.Product, *web.ErrorResponse) {
	err := service.Validate.Struct(request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		fmt.Println(&errResponse)
		return nil, &errResponse
	}
	product := domain.Product{
		Name:  request.Name,
		Price: request.Price,
	}
	product, err = service.ProductRepository.UpdateProduct(product, productID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return &product, nil
}

func (service *ProductServiceImpl) DeleteProduct(productID primitive.ObjectID) *web.ErrorResponse {
	err := service.ProductRepository.DeleteProduct(productID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return &errResponse
	}
	return nil
}
