package services

import (
	"os"

	"github.com/karlozz157/stori/app/mailers"
	"github.com/karlozz157/stori/app/models"
	"github.com/karlozz157/stori/app/parsers"
	"github.com/karlozz157/stori/app/readers"
	services_transactions "github.com/karlozz157/stori/app/services/transactions"
)

type SummaryService struct {
	mailerService      mailers.Mailer
	transactionService services_transactions.TransactionService
}

func NewSummaryService(mailerService mailers.Mailer, transactionService services_transactions.TransactionService) *SummaryService {
	return &SummaryService{
		mailerService:      mailerService,
		transactionService: transactionService,
	}
}

func (s *SummaryService) CreateSummary() error {
	transactions := s.geTransactionsFromCsv(os.Getenv("CSV_PATH"))

	if len(transactions) == 0 {
		return nil
	}

	return s.transactionService.AddTransactions(transactions)
}

func (s *SummaryService) GetSummary() models.Summary {
	balance, _ := s.transactionService.GetBalance()
	credit, _ := s.transactionService.GetCredit()
	debit, _ := s.transactionService.GetDebit()

	return models.Summary{
		Balance: balance,
		Credit:  credit,
		Debit:   debit,
	}
}

func (s *SummaryService) CreateAndSendSummary() error {
	s.CreateSummary()
	summary := s.GetSummary()
	s.sendEmail(summary)

	return nil
}

func (s *SummaryService) sendEmail(summary models.Summary) error {
	if err := s.mailerService.SendEmail(mailers.EmailData{
		To:       []string{os.Getenv("EMAIL_TO")},
		Subject:  os.Getenv("EMAIL_SUBJECT"),
		Template: os.Getenv("EMAIL_TEMPLATE"),
		Data:     summary,
	}); err != nil {
		return err
	}

	return nil
}

func (s *SummaryService) geTransactionsFromCsv(filepath string) []models.Transaction {
	var reader readers.Reader = readers.NewCsvReader(filepath)
	rows, _ := reader.GetData()

	parser := parsers.NewTransactionParser()
	transactions := parser.ParseTransactions(rows)

	return transactions
}
