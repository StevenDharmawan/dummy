package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type TransactionRequest struct {
	UserID    int                `json:"user_id" validate:"required"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id" validate:"required"`
	Quantity  int                `bson:"quantity" validate:"required"`
}
