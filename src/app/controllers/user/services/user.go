package services

import (
	transactionService "bank-app/src/app/controllers/transaction/services"
	"bank-app/src/app/controllers/user/models"
)

type UserRepositoryI interface {
	GetById(int) (*models.User, error)
	Update(*models.User) error
	Insert(*models.User) error
}

type UserService struct {
	u UserRepositoryI
	t transactionService.TransactionServiceI
}

func NewUserService(u UserRepositoryI, t transactionService.TransactionServiceI) *UserService {

	return &UserService{u, t}
}

func (userService *UserService) CreateUser(obj *models.User) error {
	err := userService.u.Insert(obj)
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserService) GetUserBalance(id int) (float64, error) {
	user, err := userService.u.GetById(id)
	if err != nil {
		return 0, err
	}
	return user.Amount, nil
}

func (userService *UserService) UpdateBalance(id int, amount float64) error {
	user, err := userService.u.GetById(id)
	if err != nil {
		return err
	}
	user.Amount = amount
	err = userService.u.Update(user)
	if err != nil {
		return err
	}

	err = userService.t.CreateTransaction(id, amount, 1)
	if err != nil {
		return err
	}
	return nil
}
