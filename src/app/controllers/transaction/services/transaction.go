package services

import "bank-app/src/app/controllers/transaction/models"

type TransactionRepositoryI interface {
	Insert(t *models.Transaction) error
}

type TransactionServiceI interface {
	CreateTransaction(userID int, amount float64, transactionType int) error
}

type TransactionService struct {
	r TransactionRepositoryI
}

func NewTransactionService(r TransactionRepositoryI) *TransactionService {
	return &TransactionService{r}
}

func (t *TransactionService) CreateTransaction(userID int, amount float64, transactionType int) error {
	_ = transactionType
	transaction := &models.Transaction{
		UserID:          userID,
		Amount:          amount,
		TransactionType: "",
		TransactionDate: "",
	}

	err := t.r.Insert(transaction)

	if err != nil {
		return err
	}

	return nil
}
