package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"phase3-gc1-shopping/helper"
	"phase3-gc1-shopping/model/web"
	"phase3-gc1-shopping/service"
)

type TransactionController interface {
	CreateTransaction(c echo.Context) error
	GetTransactionById(c echo.Context) error
	GetTransactions(c echo.Context) error
	UpdateTransaction(c echo.Context) error
	DeleteTransaction(c echo.Context) error
}

type TransactionControllerImpl struct {
	service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionControllerImpl {
	return &TransactionControllerImpl{TransactionService: transactionService}
}

func (controller *TransactionControllerImpl) CreateTransaction(c echo.Context) error {
	var request web.TransactionRequest
	err := c.Bind(&request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		return c.JSON(errResponse.Code, errResponse)
	}
	response, errResponse := controller.TransactionService.CreateTransaction(request)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}

	paymentURL := "http://localhost:8081/api/v1/payments"
	paymentRequest := web.PaymentRequest{
		TransactionID: response.ID,
		Status:        true,
	}
	errResponse = helper.CallPaymentService(paymentURL, paymentRequest)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}

	return c.JSON(http.StatusCreated, web.Response{
		Message: "success create transaction",
		Data:    response,
	})
}

func (controller *TransactionControllerImpl) GetTransactionById(c echo.Context) error {
	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errResponse := helper.ErrBadRequest(err)
		return c.JSON(errResponse.Code, errResponse)
	}
	response, errResponse := controller.TransactionService.GetTransactionById(transactionID)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusOK, web.Response{
		Message: "success get transaction",
		Data:    response,
	})
}

func (controller *TransactionControllerImpl) GetTransactions(c echo.Context) error {
	responses, errResponse := controller.TransactionService.GetTransactions()
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusOK, web.Response{
		Message: "success get all transaction",
		Data:    responses,
	})
}

func (controller *TransactionControllerImpl) UpdateTransaction(c echo.Context) error {
	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errResponse := helper.ErrBadRequest(err)
		return c.JSON(errResponse.Code, errResponse)
	}
	var request web.TransactionRequest
	err = c.Bind(&request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		return c.JSON(errResponse.Code, errResponse)
	}
	transaction, errResponse := controller.TransactionService.UpdateTransaction(request, transactionID)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusOK, web.Response{
		Message: "success update transaction",
		Data:    transaction,
	})
}

func (controller *TransactionControllerImpl) DeleteTransaction(c echo.Context) error {
	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errResponse := helper.ErrBadRequest(err)
		return c.JSON(errResponse.Code, errResponse)
	}
	errResponse := controller.TransactionService.DeleteTransaction(transactionID)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusOK, web.Response{
		Message: "success delete transaction",
	})
}
