package userHandler

import (
	"bank-app/src/app/controllers/user/services"
	"fmt"
	"net/http"
)

type UserHandler struct {
	s *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{s}
}

func (userHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Print(r.Body)
}

func (userHandler *UserHandler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User handler"))
}

func (userHandler *UserHandler) TpUpBalance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User handler"))
}
