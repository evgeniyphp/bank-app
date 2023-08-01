package userHandler

import (
	userModel "bank-app/src/app/controllers/user/models"
	"bank-app/src/app/controllers/user/services"
	"encoding/json"
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

	var data struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// TODO: validation...

	user := &userModel.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	fmt.Println(user)
	err = userHandler.s.CreateUser(user)
	if err != nil {
		http.Error(w, "Cannot create user", http.StatusBadRequest)
		return
	}

	result, err := json.Marshal(user)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	i, err := w.Write(result)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func (userHandler *UserHandler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User handler"))
}

func (userHandler *UserHandler) TpUpBalance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User handler"))
}
