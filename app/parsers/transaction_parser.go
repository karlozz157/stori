package parsers

import (
	"strconv"
	"time"

	"github.com/karlozz157/stori/app/models"
)

type transactionParser struct{}

func NewTransactionParser() *transactionParser {
	return &transactionParser{}
}

func (t *transactionParser) ParseTransactions(rows [][]string) []models.Transaction {
	var transactions []models.Transaction

	for _, row := range rows {
		transaction, err := t.parseTransaction(row)

		if err != nil {
			continue
		}

		transactions = append(transactions, *transaction)
	}

	return transactions
}

func (t *transactionParser) parseTransaction(row []string) (*models.Transaction, error) {
	date, err := time.Parse("1/2", row[1])
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return nil, err
	}

	var movement models.Movement = models.Debit
	if amount < 0 {
		movement = models.Credit
	}

	return models.NewTransaction(date, amount, movement), nil
}
