package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"phase3-gc1-payment/config"
	"phase3-gc1-payment/controller"
	"phase3-gc1-payment/repository"
	"phase3-gc1-payment/service"
)

func main() {
	ctx := context.Background()
	client := config.ConnectMongoDB(ctx)
	paymentCollection := config.GetCollection(client, "payment")

	paymentRepository := repository.NewPaymentRepository(ctx, paymentCollection)
	paymentService := service.NewPaymentService(paymentRepository)
	paymentController := controller.NewPaymentController(paymentService)

	e := echo.New()
	v1 := e.Group("/api/v1")
	{
		v1.POST("/payments", paymentController.CreatePayment)
	}
	e.Logger.Fatal(e.Start(":8081"))
}
