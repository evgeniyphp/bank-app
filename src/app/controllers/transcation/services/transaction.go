package transactionService

import transactionModel "bank-app/src/app/controllers/transcation/models"

type TransactionServiceI interface {
	CreateTransaction(userID int, amount float64, transactionType int) error
}

type TransactionService struct {
	r *transactionModel.TransactionI
}

func (t *TransactionService) CreateTransaction(userID int, amount float64, transactionType int) error {
	_ = transactionType
	transaction := &transactionModel.Transaction{
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
