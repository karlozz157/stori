package models

type Summary struct {
	Balance float64 `json:"balance"`
	Credit  float64 `json:"credit"`
	Debit   float64 `json:"debit"`
}
