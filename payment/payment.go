package payment

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreatePayment(c echo.Context) error {
	return c.String(http.StatusOK, "create payment")
}

func UpdateTransactionStatus(c echo.Context) error {
	return c.String(http.StatusOK, "update payment transaction status")
}

func GetTransactionStatus(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("got id:", id)
	return c.String(http.StatusOK, fmt.Sprintf("update transaction status by id %s", id))
}
