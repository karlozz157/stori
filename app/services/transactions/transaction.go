package services_transactions

import (
	"time"

	"github.com/karlozz157/stori/app/models"
)

type TransactionService interface {
	AddTransaction(transaction models.Transaction) error
	AddTransactions(transactions []models.Transaction) error
	GetBalance() (float64, error)
	GetCredit() (float64, error)
	GetDebit() (float64, error)
	GetNumberOfTransactionsGrouped() (map[time.Month]int, error)
	Init()
}
