package services_transactions_adapters_test

import (
	"testing"
	"time"

	"github.com/karlozz157/stori/app/models"
	services_transactions_adapters "github.com/karlozz157/stori/app/services/transactions/adapters"
)

func TestTransactionStorageService(t *testing.T) {

	transactions := []models.Transaction{
		{
			Id:       1,
			Date:     time.Now(),
			Amount:   2200,
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
			Amount:   +300,
			Movement: models.Debit,
		},
	}

	service := services_transactions_adapters.NewTransactionStorageService()
	service.Init()

	for _, transaction := range transactions {
		if err := service.AddTransaction(transaction); err != nil {
			t.Errorf("Insertado %s", err.Error())
		}
	}

	balance, _ := service.GetBalance()
	expectedBalance := -1600.00
	if balance != expectedBalance {
		t.Errorf("Balance incorrecto. Se esperaba %.2f pero se obtuvo %.2f", expectedBalance, balance)
	}

	credit, _ := service.GetCredit()
	expectedCredit := 2200.00
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
