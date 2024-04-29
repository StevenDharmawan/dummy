package service

import (
	"phase3-gc1-payment/helper"
	"phase3-gc1-payment/model/domain"
	"phase3-gc1-payment/model/web"
	"phase3-gc1-payment/repository"
)

type PaymentService interface {
	CreatePayment(request web.PaymentRequest) *web.ErrorResponse
}

type PaymentServiceImpl struct {
	repository.PaymentRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository) *PaymentServiceImpl {
	return &PaymentServiceImpl{PaymentRepository: paymentRepository}
}

func (service *PaymentServiceImpl) CreatePayment(request web.PaymentRequest) *web.ErrorResponse {
	payment := domain.Payment{
		TransactionID: request.TransactionID,
		Status:        request.Status,
	}
	err := service.PaymentRepository.CreatePayment(payment)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return &errResponse
	}
	return nil
}
