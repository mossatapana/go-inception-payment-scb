package main

import (
	"go-inception-payment-scb/payment"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/payment", payment.CreatePayment)
	e.PUT("/payment/status", payment.UpdateTransactionStatus)
	e.GET("/transaction/:id/status", payment.GetTransactionStatus)
	e.Logger.Fatal(e.Start(":8080"))
}
