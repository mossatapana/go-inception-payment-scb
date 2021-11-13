package repository

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	"go-inception-payment-scb/model"
)

type PaymentRepository interface {
	Insert(data model.PaymentORM) (int64, error)
	UpdateStatus(id int, status string) error
	GetByID(id int) (model.PaymentORM, error)
}

type PaymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (pr PaymentRepo) Insert(data model.PaymentORM) (int64, error) {
	statement, err := pr.db.Prepare(`INSERT INTO payment(
		"amount",
		"card",
		"currency",
		"status",
		"capture",
		"authorized",
		"reversed",
		"paid",
		"transaction",
		"offsite",
		"created_at",
		"updated_at"
	) values (?,?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		log.Errorf("prepare insert statement error %v\n", err)
		return 0, err
	}

	res, err := statement.Exec(data.Amount, data.Card, data.Currency, data.Status, data.Capture,
		data.Authorized, data.Reversed, data.Paid, data.Transaction, data.OffsiteType, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		log.Errorf("insert payment error %v\n", err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Errorf("get last payment id record error %v\n", err)
		return 0, err
	}
	return id, nil
}

func (pr PaymentRepo) UpdateStatus(id int, status string) error {
	return nil
}

func (pr PaymentRepo) GetByID(id int) (model.PaymentORM, error) {
	var payment model.PaymentORM
	return payment, nil
}
