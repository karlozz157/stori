package services

import (
	"os"

	"github.com/karlozz157/stori/app/mailers"
	mailers_adapaters "github.com/karlozz157/stori/app/mailers/adapters"
	"github.com/karlozz157/stori/app/models"
	"github.com/karlozz157/stori/app/parsers"
	"github.com/karlozz157/stori/app/readers"
	services_transactions "github.com/karlozz157/stori/app/services/transactions"
	services_transactions_adapters "github.com/karlozz157/stori/app/services/transactions/adapters"
)

type summary struct {
}

func NewSummaryService() *summary {
	return &summary{}
}

func (s *summary) SendSummaryEmail() error {
	transactions := s.geTransactionsFromCsv(os.Getenv("CSV_PATH"))

	var transactionService services_transactions.TransactionService = services_transactions_adapters.NewTransactionMemoryService() // or u can use adapters.NewTransactionMemoryService()

	if len(transactions) > 0 {
		transactionService.AddTransactions(transactions)
	}

	s.sendEmail(transactionService)

	return nil
}

func (s *summary) geTransactionsFromCsv(filepath string) []models.Transaction {
	var reader readers.Reader = readers.NewCsvReader(filepath)
	rows, _ := reader.GetData()

	parser := parsers.NewTransactionParser()
	transactions := parser.ParseTransactions(rows)

	return transactions
}

func (s *summary) sendEmail(transactionService services_transactions.TransactionService) error {
	var mailer mailers.Mailer = mailers_adapaters.NewEmailService()

	balance, _ := transactionService.GetBalance()
	credit, _ := transactionService.GetCredit()
	debit, _ := transactionService.GetDebit()

	data := struct {
		Balance float64
		Credit  float64
		Debit   float64
	}{
		Balance: balance,
		Credit:  credit,
		Debit:   debit,
	}

	return mailer.SendEmail(mailers.EmailData{
		To:       []string{os.Getenv("EMAIL_TO")},
		Subject:  os.Getenv("EMAIL_SUBJECT"),
		Template: os.Getenv("EMAIL_TEMPLATE"),
		Data:     data,
	})
}
