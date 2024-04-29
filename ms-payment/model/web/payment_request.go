package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentRequest struct {
	TransactionID primitive.ObjectID `json:"transaction_id"`
	Status        bool               `json:"status"`
}
