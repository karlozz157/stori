package adapters_test

import (
	"testing"
	"time"

	"github.com/karlozz157/stori/app/models"
	"github.com/karlozz157/stori/app/parsers"
)

func TestTransactionParser(t *testing.T) {
	rows := [][]string{
		{"0", "7/15", "+60.5"},
		{"1", "7/28", "-10.3"},
		{"2", "8/2", "-20.46"},
		{"3", "8/12", "+10"},
	}

	transactionParser := parsers.NewTransactionParser()
	transactions := transactionParser.ParseTransactions(rows)

	firstTransaction := transactions[0]

	expectedAmount := 60.5
	if firstTransaction.Amount != expectedAmount {
		t.Errorf("Amount incorrecto. Se esperaba %.2f", expectedAmount)
	}

	if firstTransaction.Movement != models.Debit {
		t.Errorf("Movimiento incorrecto. Se esperaba %s", models.Debit)
	}

	if firstTransaction.Date.Month() != time.July {
		t.Errorf("Mes incorrecto. Se esperaba %s", time.July)
	}
}
