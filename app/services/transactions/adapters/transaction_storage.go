package services_transactions_adapters

import (
	"database/sql"
	"math"
	"time"

	"github.com/karlozz157/stori/app/models"
	services_transactions "github.com/karlozz157/stori/app/services/transactions"
	_ "github.com/mattn/go-sqlite3"
)

type transactionStorage struct {
	db *sql.DB
}

func NewTransactionStorageService() services_transactions.TransactionService {
	db, err := sql.Open("sqlite3", "./stori.db")
	if err != nil {
		panic(err)
	}

	return &transactionStorage{
		db: db,
	}
}

func (t *transactionStorage) AddTransaction(transaction models.Transaction) error {
	transaction.Amount = math.Abs(transaction.Amount)

	sql := "INSERT INTO transactions(amount, date, movement) VALUES(?, ?, ?)"
	_, err := t.db.Exec(sql, transaction.Amount, transaction.Date, transaction.Movement)

	return err
}

func (t *transactionStorage) AddTransactions(transactions []models.Transaction) error {
	for _, transaction := range transactions {
		if err := t.AddTransaction(transaction); err != nil {
			return err
		}
	}

	return nil
}

func (t *transactionStorage) GetBalance() (float64, error) {
	var balance float64

	credit, err := t.getBalanceByMovement(models.Credit)
	if err != nil {
		return balance, err
	}

	debit, err := t.getBalanceByMovement(models.Debit)
	if err != nil {
		return balance, err
	}

	return (debit - credit), nil
}

func (t *transactionStorage) GetCredit() (float64, error) {
	return t.getBalanceByMovement(models.Credit)
}

func (t *transactionStorage) GetDebit() (float64, error) {
	return t.getBalanceByMovement(models.Debit)
}

func (t *transactionStorage) GetNumberOfTransactionsGrouped() (map[time.Month]int, error) {
	grouped := make(map[time.Month]int)

	sql := "SELECT date FROM transactions"

	rows, err := t.db.Query(sql)
	if err != nil {
		return grouped, err
	}

	for rows.Next() {
		var date time.Time
		if err := rows.Scan(&date); err != nil {
			return nil, err
		}

		month := date.Month()
		grouped[month]++
	}

	return grouped, nil
}

func (t *transactionStorage) Init() {
	t.createTable()
	t.truncateTable()
}

func (t *transactionStorage) createTable() {
	_, err := t.db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date TIMESTAMP NOT NULL,
			amount REAL NOT NULL,
			movement TEXT NOT NULL
		)`)

	if err != nil {
		panic(err)
	}
}

func (t *transactionStorage) getBalanceByMovement(movement models.Movement) (float64, error) {
	var balance float64

	sql := "SELECT SUM(amount) FROM transactions WHERE movement = ?"
	err := t.db.QueryRow(sql, movement).Scan(&balance)

	return balance, err
}

func (t *transactionStorage) truncateTable() {
	_, err := t.db.Exec(`DELETE FROM transactions`)

	if err != nil {
		panic(err)
	}
}
