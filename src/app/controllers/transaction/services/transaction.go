package services

import "bank-app/src/app/controllers/transaction/models"

type transactionRepositoryI interface {
	Insert(t *models.Transaction) error
}

type TransactionService struct {
	r transactionRepositoryI
}

func NewTransactionService(r transactionRepositoryI) *TransactionService {
	return &TransactionService{r}
}

func (t *TransactionService) CreateTransaction(userID int, amount float64, transactionType int) error {
	_ = transactionType
	transaction := &models.Transaction{
		UserID:          userID,
		Amount:          amount,
		TransactionType: "", // need to add enum to determine type of transaction
		TransactionDate: "",
	}

	err := t.r.Insert(transaction)

	if err != nil {
		return err
	}

	return nil
}
