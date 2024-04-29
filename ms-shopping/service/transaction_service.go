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

type TransactionService interface {
	CreateTransaction(request web.TransactionRequest) (*domain.Transaction, *web.ErrorResponse)
	GetTransactions() ([]domain.Transaction, *web.ErrorResponse)
	GetTransactionById(transactionID primitive.ObjectID) (*domain.Transaction, *web.ErrorResponse)
	UpdateTransaction(request web.TransactionRequest, transactionID primitive.ObjectID) (*domain.Transaction, *web.ErrorResponse)
	DeleteTransaction(transactionID primitive.ObjectID) *web.ErrorResponse
}

type TransactionServiceImpl struct {
	repository.TransactionRepository
	repository.ProductRepository
	*validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, productRepository repository.ProductRepository, validate *validator.Validate) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		ProductRepository:     productRepository,
		Validate:              validate,
	}
}

func (service *TransactionServiceImpl) CreateTransaction(request web.TransactionRequest) (*domain.Transaction, *web.ErrorResponse) {
	err := service.Validate.Struct(request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		fmt.Println(&errResponse)
		return nil, &errResponse
	}
	product, err := service.ProductRepository.GetProductById(request.ProductID)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		return nil, &errResponse
	}
	transaction := domain.Transaction{
		UserID:     request.UserID,
		ProductID:  request.ProductID,
		Quantity:   request.Quantity,
		TotalPrice: float64(request.Quantity) * product.Price,
	}
	transaction, err = service.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return &transaction, nil
}

func (service *TransactionServiceImpl) GetTransactions() ([]domain.Transaction, *web.ErrorResponse) {
	transactions, err := service.TransactionRepository.GetAllTransactions()
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return transactions, nil
}

func (service *TransactionServiceImpl) GetTransactionById(transactionID primitive.ObjectID) (*domain.Transaction, *web.ErrorResponse) {
	transaction, err := service.TransactionRepository.GetTransactionById(transactionID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return &transaction, nil
}

func (service *TransactionServiceImpl) UpdateTransaction(request web.TransactionRequest, transactionID primitive.ObjectID) (*domain.Transaction, *web.ErrorResponse) {
	err := service.Validate.Struct(request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		fmt.Println(&errResponse)
		return nil, &errResponse
	}
	product, err := service.ProductRepository.GetProductById(request.ProductID)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		return nil, &errResponse
	}
	transaction := domain.Transaction{
		UserID:     request.UserID,
		ProductID:  request.ProductID,
		Quantity:   request.Quantity,
		TotalPrice: float64(request.Quantity) * product.Price,
	}
	transaction, err = service.TransactionRepository.UpdateTransaction(transaction, transactionID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return &transaction, nil
}

func (service *TransactionServiceImpl) DeleteTransaction(transactionID primitive.ObjectID) *web.ErrorResponse {
	err := service.TransactionRepository.DeleteTransaction(transactionID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return &errResponse
	}
	return nil
}
