package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"phase3-gc1-shopping/config"
	"phase3-gc1-shopping/controller"
	"phase3-gc1-shopping/repository"
	"phase3-gc1-shopping/service"
)

func main() {
	ctx := context.Background()
	validate := validator.New()
	client := config.ConnectMongoDB(ctx)
	transactionCollection := config.GetCollection(client, "transactions")
	productCollection := config.GetCollection(client, "product")

	productRepository := repository.NewProductRepository(ctx, productCollection)
	productService := service.NewProductService(productRepository, validate)
	productController := controller.NewProductController(productService)

	transactionRepository := repository.NewTransactionRepository(ctx, transactionCollection)
	transactionService := service.NewTransactionService(transactionRepository, productRepository, validate)
	transactionController := controller.NewTransactionController(transactionService)

	e := echo.New()
	v1 := e.Group("/api/v1")
	{
		v1.POST("/transactions", transactionController.CreateTransaction)
		v1.GET("/transactions", transactionController.GetTransactions)
		v1.GET("/transactions/:id", transactionController.GetTransactionById)
		v1.PUT("/transactions/:id", transactionController.UpdateTransaction)
		v1.DELETE("/transactions/:id", transactionController.DeleteTransaction)

		v1.POST("/products", productController.CreateProduct)
		v1.GET("/products", productController.GetProducts)
		v1.GET("/products/:id", productController.GetProductById)
		v1.PUT("/products/:id", productController.UpdateProduct)
		v1.DELETE("/products/:id", productController.DeleteProduct)
		v1.DELETE("/products/:id", productController.DeleteProduct)

	}
	e.Logger.Fatal(e.Start(":8080"))
}
