package payment

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"go-inception-payment-scb/model"
	"go-inception-payment-scb/repository"
	"net/http"
	"time"
)

type Payment interface {
	CreatePayment(c echo.Context) error
	UpdateTransactionStatus(c echo.Context) error
	GetTransactionStatus(c echo.Context) error
}

type PaymentController struct {
	client      *omise.Client
	paymentRepo *repository.PaymentRepo
}

func NewPayment(client *omise.Client, paymentRepo *repository.PaymentRepo) *PaymentController {
	return &PaymentController{client: client, paymentRepo: paymentRepo}
}

func (p PaymentController) CreatePayment(c echo.Context) error {
	res := model.CreatePaymentResponse{}
	tn := time.Now()
	paymentRequest := model.CreatePaymentRequest{}
	if err := c.Bind(&paymentRequest); err != nil {
		log.Errorf("bind req error: %v\n", err)
		res.Message = "request error"
		return c.JSON(http.StatusBadRequest, res)
	}

	createToken := operations.CreateToken{
		Name:            paymentRequest.Name,
		Number:          paymentRequest.Number,
		ExpirationMonth: paymentRequest.ExpirationMonth,
		ExpirationYear:  paymentRequest.ExpirationYear,
		SecurityCode:    paymentRequest.SecurityCode,
		City:            paymentRequest.City,
		PostalCode:      paymentRequest.PostalCode,
	}
	token, err := p.createToken(&createToken)
	if err != nil {
		log.Errorf("create token error: %v\n", err)
		res.Message = "unable to processed request"
		return c.JSON(http.StatusInternalServerError, res)
	}

	createCharge := &operations.CreateCharge{
		Amount:   paymentRequest.Amount,
		Card:     token.ID,
		Currency: "thb",
	}

	var charge omise.Charge
	if err := p.client.Do(&charge, createCharge); err != nil {
		log.Errorf("create token error: %v\n", err)
		res.Message = "unable to processed request"
		return c.JSON(http.StatusInternalServerError, res)
	}

	data := model.PaymentORM{
		Amount:      int(charge.Amount),
		Card:        charge.Card.ID,
		Currency:    charge.Currency,
		Status:      string(charge.Status),
		Capture:     charge.Capture,
		Authorized:  charge.Authorized,
		Reversed:    charge.Reversed,
		Paid:        charge.Paid,
		Transaction: charge.Transaction,
		OffsiteType: string(charge.Offsite),
		CreatedAt:   tn,
		UpdatedAt:   tn,
	}

	if charge.Description != nil {
		data.Description = *charge.Description
	} else {
		data.Description = ""
	}

	id, err := p.paymentRepo.Insert(data)
	if err != nil {
		log.Errorf("insert payment record error: %v\n", err)
		res.Message = "unable to processed request"
		return c.JSON(http.StatusServiceUnavailable, res)
	}

	res.Message = "success"
	res.ID = id
	return c.JSON(http.StatusOK, res)
}

func (p PaymentController) UpdateTransactionStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, "update payment transaction status")
}

func (p PaymentController) GetTransactionStatus(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("got id:", id)
	return c.JSON(http.StatusOK, fmt.Sprintf("update transaction status by id %s", id))
}

func (p PaymentController) createToken(createToken *operations.CreateToken) (*omise.Token, error) {
	token := &omise.Token{}
	if err := p.client.Do(token, createToken); err != nil {
		return token, fmt.Errorf("create token for create payment error: %v", err)
	}
	fmt.Println("created token:", token.ID)
	return token, nil
}
