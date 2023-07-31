package services

import userModel "bank-app/src/app/controllers/user/models"

type UserService struct {
	u *userModel.UserRepository
}

func NewUserService(userReposutory *userModel.UserRepository) *UserService {
	return &UserService{userReposutory}
}

func (userService *UserService) CreateUser(obj *userModel.User) error {
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
	return nil
}
