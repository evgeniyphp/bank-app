package models

import (
	"bank-app/pkg/database/sqlite3"

	"testing"

	"github.com/stretchr/testify/assert"
)


func TestInsert(t *testing.T) {
	storage, err := sqlite3.New(":memory:")
	db := storage.DB

	defer db.Close()

	if err != nil {
		t.Fatal("Cannot create db for tests")
	}

	repository := NewTransactionRepository(db)

	transaction := Transaction{
		UserID:          1,
		Amount:          100.0,
		TransactionDate: "2023-09-05",
		TransactionType: "Credit",
	}

	err = repository.Insert(&transaction)

	assert.NoError(t, err)
}
