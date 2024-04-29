package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"phase3-gc1-shopping/model/domain"
)

type TransactionRepository interface {
	CreateTransaction(transaction domain.Transaction) (domain.Transaction, error)
	GetAllTransactions() ([]domain.Transaction, error)
	GetTransactionById(transactionID primitive.ObjectID) (domain.Transaction, error)
	UpdateTransaction(transaction domain.Transaction, transactionID primitive.ObjectID) (domain.Transaction, error)
	DeleteTransaction(transactionID primitive.ObjectID) error
}

type TransactionRepositoryImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewTransactionRepository(ctx context.Context, collection *mongo.Collection) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{ctx: ctx, collection: collection}
}

func (repository *TransactionRepositoryImpl) CreateTransaction(transaction domain.Transaction) (domain.Transaction, error) {
	result, err := repository.collection.InsertOne(repository.ctx, transaction)
	transaction.ID = result.InsertedID.(primitive.ObjectID)
	return transaction, err
}

func (repository *TransactionRepositoryImpl) GetAllTransactions() ([]domain.Transaction, error) {
	cursor, err := repository.collection.Find(repository.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(repository.ctx)
	var transactions []domain.Transaction
	for cursor.Next(repository.ctx) {
		var transaction domain.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (repository *TransactionRepositoryImpl) GetTransactionById(transactionID primitive.ObjectID) (domain.Transaction, error) {
	var transaction domain.Transaction
	filter := bson.M{"_id": transactionID}
	err := repository.collection.FindOne(repository.ctx, filter).Decode(&transaction)
	return transaction, err
}

func (repository *TransactionRepositoryImpl) UpdateTransaction(transaction domain.Transaction, transactionID primitive.ObjectID) (domain.Transaction, error) {
	filter := bson.M{"_id": transactionID}
	update := bson.M{
		"$set": transaction,
	}
	_, err := repository.collection.UpdateOne(context.Background(), filter, update)
	transaction.ID = transactionID
	return transaction, err
}

func (repository *TransactionRepositoryImpl) DeleteTransaction(transactionID primitive.ObjectID) error {
	_, err := repository.collection.DeleteOne(repository.ctx, bson.M{"_id": transactionID})
	return err
}
