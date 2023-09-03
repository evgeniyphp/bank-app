package services

import (
	"bank-app/src/app/controllers/user/models"
)

type userRepositoryI interface {
	GetById(int) (*models.User, error)
	Update(*models.User) error
	Insert(*models.User) error
}

type transactionServiceI interface {
	CreateTransaction(userID int, amount float64, transactionType int) error
}

type UserService struct {
	u userRepositoryI
	t transactionServiceI
}

func NewUserService(u userRepositoryI, t transactionServiceI) *UserService {

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
