package services_transactions_adapters

import (
	"math"
	"time"

	"github.com/karlozz157/stori/app/models"
	services_transactions "github.com/karlozz157/stori/app/services/transactions"
)

type transactionMemory struct {
	transactions []models.Transaction
}

func NewTransactionMemoryService() services_transactions.TransactionService {
	return &transactionMemory{}
}

func (t *transactionMemory) AddTransaction(transaction models.Transaction) error {
	transaction.Amount = math.Abs(transaction.Amount)
	t.transactions = append(t.transactions, transaction)

	return nil
}

func (t *transactionMemory) AddTransactions(transactions []models.Transaction) error {
	for _, transaction := range transactions {
		t.AddTransaction(transaction)
	}

	return nil
}

func (t *transactionMemory) GetBalance() (float64, error) {
	var balance float64

	for _, transaction := range t.transactions {

		if transaction.Movement == models.Debit {
			balance += transaction.Amount
		} else {
			balance -= transaction.Amount
		}
	}

	return balance, nil
}

func (t *transactionMemory) GetAverageCredit() (float64, error) {
	return t.getAverageByMovement(models.Credit)
}

func (t *transactionMemory) GetAverageDebit() (float64, error) {
	return t.getAverageByMovement(models.Debit)
}

func (t *transactionMemory) GetNumberOfTransactionsGrouped() (map[time.Month]int, error) {
	grouped := make(map[time.Month]int)

	for _, transaction := range t.transactions {
		month := transaction.Date.Month()
		grouped[month] += 1
	}

	return grouped, nil
}

func (t *transactionMemory) Init() {
	t.transactions = []models.Transaction{}
}

func (t *transactionMemory) getAverageByMovement(movement models.Movement) (float64, error) {
	var balance, result float64
	movements := 0

	for _, transaction := range t.transactions {
		if transaction.Movement == movement {
			movements++
			balance += transaction.Amount
		}
	}

	if movements > 0 {
		result = balance / float64(movements)
	}

	return result, nil
}
