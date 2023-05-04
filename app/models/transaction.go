package models

import (
	"math"
	"time"
)

type Movement string

const (
	Debit  Movement = "debit"
	Credit Movement = "credit"
)

type Transaction struct {
	Id       int
	Date     time.Time
	Amount   float64
	Movement Movement
}

func NewTransaction(date time.Time, amount float64, movement Movement) *Transaction {
	amount = math.Abs(amount)

	return &Transaction{
		Date:     date,
		Amount:   amount,
		Movement: movement,
	}
}
