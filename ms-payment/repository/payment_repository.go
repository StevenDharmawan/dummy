package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"phase3-gc1-payment/model/domain"
)

type PaymentRepository interface {
	CreatePayment(product domain.Payment) error
}

type PaymentRepositoryImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewPaymentRepository(ctx context.Context, collection *mongo.Collection) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{ctx: ctx, collection: collection}
}

func (repository *PaymentRepositoryImpl) CreatePayment(product domain.Payment) error {
	_, err := repository.collection.InsertOne(repository.ctx, product)
	return err
}
