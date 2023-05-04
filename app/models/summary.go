package models

import "time"

type Summary struct {
	Balance              float64            `json:"balance"`
	Credit               float64            `json:"credit"`
	Debit                float64            `json:"debit"`
	NumberOfTransactions map[time.Month]int `json:"number_of_transactions"`
}
