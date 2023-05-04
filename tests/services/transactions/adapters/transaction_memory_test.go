package services_transactions_adapters_test

import (
	"testing"
	"time"

	"github.com/karlozz157/stori/app/models"
	services_transactions_adapters "github.com/karlozz157/stori/app/services/transactions/adapters"
)

func TestTransactionMemoryService(t *testing.T) {

	transactions := []models.Transaction{
		{
			Id:       1,
			Date:     time.Now(),
			Amount:   -1000,
			Movement: models.Credit,
		},
		{
			Id:       2,
			Date:     time.Now(),
			Amount:   300,
			Movement: models.Debit,
		},
		{
			Id:       3,
			Date:     time.Now().AddDate(0, 3, 0),
			Amount:   300,
			Movement: models.Debit,
		},
	}

	service := services_transactions_adapters.NewTransactionMemoryService()

	for _, transaction := range transactions {
		service.AddTransaction(transaction)
	}

	balance, _ := service.GetBalance()
	expectedBalance := -400.00
	if balance != expectedBalance {
		t.Errorf("Balance incorrecto. Se esperaba %.2f pero se obtuvo %.2f", expectedBalance, balance)
	}

	credit, _ := service.GetCredit()
	expectedCredit := 1000.00
	if credit != expectedCredit {
		t.Errorf("Credito incorrecto. Se esperaba %.2f pero se obtuvo %.2f", expectedCredit, credit)

	}

	debit, _ := service.GetDebit()
	expectedDebit := 600.00
	if debit != expectedDebit {
		t.Errorf("Debito incorrecto. Se esperaba %.2f pero se obtuvo %.2f", expectedDebit, debit)
	}

	month := time.Now().Month()
	grouped, _ := service.GetNumberOfTransactionsGrouped()

	numberOfTransaction := grouped[month]
	expectedNumberOfTransaction := 2

	if numberOfTransaction != expectedNumberOfTransaction {
		t.Errorf("Número de transacciones incorrectos. Se esperaba %d pero se obtuvo %d", expectedNumberOfTransaction, numberOfTransaction)
	}
}
