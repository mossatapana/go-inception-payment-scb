package model

import "time"

type CreatePaymentRequest struct {
	Name            string     `json:"name"`
	Number          string     `json:"number"`
	ExpirationMonth time.Month `json:"expiration_month"`
	ExpirationYear  int        `json:"expiration_year"`
	SecurityCode    string     `json:"security_code"`
	City            string     `json:"city,omitempty"`
	PostalCode      string     `json:"postal_code,omitempty"`
	Amount          int64      `json:"amount"`
}

type CreatePaymentResponse struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message"`
}

type PaymentORM struct {
	ID          int       `json:"id"`
	Amount      int       `json:"amount"`
	Card        string    `json:"card"`
	Currency    string    `json:"currency"`
	Status      string    `json:"status"`
	Description *string   `json:"description"`
	Capture     bool      `json:"capture"`
	Authorized  bool      `json:"authorized"`
	Reversed    bool      `json:"reversed"`
	Paid        bool      `json:"paid"`
	Transaction string    `json:"transaction"`
	OffsiteType string    `json:"offsite"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdatePaymentStatusRequest struct {
}

type GetPaymentTransactionResponse struct {
	PaymentORM []PaymentORM `json:"transaction,omitempty"`
	Message    string       `json:"message"`
}
