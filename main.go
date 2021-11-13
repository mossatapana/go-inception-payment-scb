package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/omise/omise-go"
	"go-inception-payment-scb/payment"
	"go-inception-payment-scb/repository"
	"log"
	"os"
)

func main() {
	os.Remove("sqlite-database.db")
	fmt.Println("creating sqlite-database.db file...")

	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatalf("create file sqlite-database.db error %s\n", err)
	}
	file.Close()
	fmt.Println("sqlite-database.db has been created")

	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sqlite 3 connect: %v\n", db)
	defer db.Close()
	createTable(db)

	config, err := initConfig("config")
	if err != nil {
		log.Fatalf("init config error: %v", err)

	}
	fmt.Printf("config got: %s\n", config)

	client, err := omise.NewClient(config.Omise.PublicKey, config.Omise.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("client got: %s\n", client)

	paymentRepo := repository.NewPaymentRepo(db)
	paymentCtrl := payment.NewPayment(client, paymentRepo)

	e := echo.New()
	e.POST("/payment", paymentCtrl.CreatePayment)
	e.PUT("/payment/status", paymentCtrl.UpdateTransactionStatus)
	e.GET("/transaction/:id/status", paymentCtrl.GetTransactionStatus)
	e.Logger.Fatal(e.Start(":8080"))
}

func createTable(db *sql.DB) {
	createPaymentTableSQL := `CREATE TABLE payment (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"amount" integer,
		"card" text,
		"currency" text,
		"status" text,
		"capture" boolean,
		"authorized" boolean,
		"reversed" boolean,
		"paid" boolean,
		"transaction" text,
		"offsite" text,
		"created_at" datetime,
		"updated_at" datetime
	)`
	statement, err := db.Prepare(createPaymentTableSQL)
	if err != nil {
		log.Fatalf("prepare statement to create table payment error: %v\n", err.Error())
	}
	statement.Exec()
	log.Println("payment table created")
}
