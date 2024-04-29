package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"phase3-gc1-shopping/model/domain"
)

type ProductRepository interface {
	CreateProduct(product domain.Product) (domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	GetProductById(productID primitive.ObjectID) (domain.Product, error)
	UpdateProduct(product domain.Product, productID primitive.ObjectID) (domain.Product, error)
	DeleteProduct(productID primitive.ObjectID) error
}

type ProductRepositoryImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewProductRepository(ctx context.Context, collection *mongo.Collection) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{ctx: ctx, collection: collection}
}

func (repository *ProductRepositoryImpl) CreateProduct(product domain.Product) (domain.Product, error) {
	result, err := repository.collection.InsertOne(repository.ctx, product)
	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, err
}

func (repository *ProductRepositoryImpl) GetAllProducts() ([]domain.Product, error) {
	cursor, err := repository.collection.Find(repository.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(repository.ctx)
	var products []domain.Product
	for cursor.Next(repository.ctx) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repository *ProductRepositoryImpl) GetProductById(productID primitive.ObjectID) (domain.Product, error) {
	var product domain.Product
	filter := bson.M{"_id": productID}
	err := repository.collection.FindOne(repository.ctx, filter).Decode(&product)
	return product, err
}

func (repository *ProductRepositoryImpl) UpdateProduct(product domain.Product, productID primitive.ObjectID) (domain.Product, error) {
	filter := bson.M{"_id": productID}
	update := bson.M{
		"$set": product,
	}
	_, err := repository.collection.UpdateOne(context.Background(), filter, update)
	product.ID = productID
	return product, err
}

func (repository *ProductRepositoryImpl) DeleteProduct(productID primitive.ObjectID) error {
	_, err := repository.collection.DeleteOne(repository.ctx, bson.M{"_id": productID})
	return err
}
