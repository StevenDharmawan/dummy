package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     int                `json:"user_id"`
	ProductID  primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity   int                `bson:"quantity"`
	TotalPrice float64            `bson:"totalPrice"`
}
