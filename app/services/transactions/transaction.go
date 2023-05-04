package services_transactions

import (
	"time"

	"github.com/karlozz157/stori/app/models"
)

type TransactionService interface {
	AddTransaction(transaction models.Transaction) error
	AddTransactions(transactions []models.Transaction) error
	GetBalance() (float64, error)
	GetAverageCredit() (float64, error)
	GetAverageDebit() (float64, error)
	GetNumberOfTransactionsGrouped() (map[time.Month]int, error)
	Init()
}
