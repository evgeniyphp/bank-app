package services

import (
	purchaseModel "bank-app/src/app/controllers/purchase/models"
	transactionModel "bank-app/src/app/controllers/transaction/models"
	userModel "bank-app/src/app/controllers/user/models"
)

type goodRepositoryI interface {
	Insert(*purchaseModel.Good) error
	GetById(int) (*purchaseModel.Good, error)
}

type userRepositoryI interface {
	GetById(int) (*userModel.User, error)
	Update(*userModel.User) error
}

type transactionI interface {
	Insert(*transactionModel.Transaction) error
}

type PurchaseService struct {
	r goodRepositoryI
	u userRepositoryI
	t transactionI
}

func NewPurchaseService(r goodRepositoryI, u userRepositoryI, t transactionI) *PurchaseService {
	return &PurchaseService{r, u, t}
}

func (p *PurchaseService) CreateGood(good *purchaseModel.Good) error {
	err := p.r.Insert(good)
	if err != nil {
		return err
	}
	return nil
}

// The method to buy a good. It takes userID and goodID
// Return error. Could return custom error NotEnoughBalance
// in case if user doesn't have enough balance
func (p *PurchaseService) BuyGood(userID int, goodID int) error {
	user, err := p.u.GetById(userID)
	if err != nil {
		return err
	}

	good, err := p.r.GetById(goodID)
	if err != nil {
		return err
	}

	if user.Amount < good.Price {
		return &NotEnoughBalance{}
	}

	user.Amount = user.Amount - good.Price
	err = p.u.Update(user)
	if err != nil {
		return err
	}

	transaction := &transactionModel.Transaction{
		UserID:          userID,
		Amount:          good.Price,
		TransactionDate: "",
		TransactionType: "2", // transaction type. need to add enum
	}
	err = p.t.Insert(transaction)
	if err != nil {
		return err
	}

	return nil
}

type NotEnoughBalance struct{}

func (e *NotEnoughBalance) Error() string {
	return "Not enoguh balance"
}
