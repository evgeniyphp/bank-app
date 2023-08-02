package purchaseService

import (
	purchaseModel "bank-app/src/app/controllers/purchase/models"
	transactionModel "bank-app/src/app/controllers/transcation/models"
	user_model "bank-app/src/app/controllers/user/models"
)

type PurchaseServiceI interface {
	CreateGood(*purchaseModel.Good) error
	BuyGood(int, int) error
}

type PurchaseService struct {
	r purchaseModel.GoodRepositoryI
	u user_model.UserRepositoryI
	t transactionModel.TransactionI
}

func NewPurhcaseService(r purchaseModel.GoodRepositoryI, u user_model.UserRepositoryI, t transactionModel.TransactionI) *PurchaseService {
	return &PurchaseService{r, u, t}
}

func (p *PurchaseService) CreateGood(good *purchaseModel.Good) error {
	err := p.r.Insert(good)
	if err != nil {
		return err
	}
	return nil
}

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
		TransactionType: "2",
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
