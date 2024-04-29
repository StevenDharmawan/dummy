package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"phase3-gc1-payment/helper"
	"phase3-gc1-payment/model/web"
	"phase3-gc1-payment/service"
)

type PaymentController interface {
	CreatePayment(c echo.Context) error
}

type PaymentControllerImpl struct {
	service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) *PaymentControllerImpl {
	return &PaymentControllerImpl{PaymentService: paymentService}
}

func (controller *PaymentControllerImpl) CreatePayment(c echo.Context) error {
	var request web.PaymentRequest
	err := c.Bind(&request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		return c.JSON(errResponse.Code, errResponse)
	}
	errResponse := controller.PaymentService.CreatePayment(request)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusCreated, web.Response{
		Message: "success create product",
	})
}
