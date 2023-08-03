package transactionModel

import (
	"database/sql"
	"fmt"
)

type TransactionI interface {
	Insert(*Transaction) error
}

type Transaction struct {
	ID              int
	UserID          int
	Amount          float64
	TransactionDate string
	TransactionType string
}

type TransactionRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) Insert(t *Transaction) error {
	stmt, err := r.db.Prepare(`INSERT INTO transactions(user_id, amount, transaction_date, transaction_type) VALUES(?,?,?,?)`)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(t.UserID, t.Amount, t.TransactionDate, t.TransactionType)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
