package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TransactionID primitive.ObjectID `json:"transaction_id"`
	Status        bool               `json:"status"`
}
