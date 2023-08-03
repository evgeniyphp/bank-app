package userHandler

import (
	userModel "bank-app/src/app/controllers/user/models"
	userService "bank-app/src/app/controllers/user/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	s *userService.UserService
}

func NewUserHandler(s *userService.UserService) *UserHandler {
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

	err = userHandler.s.CreateUser(user)
	// TODO: return user to display
	if err != nil {
		http.Error(w, "Cannot create user", http.StatusBadRequest)
		return
	}

	//	result, err := json.Marshal(user)

	result := `{"success": True}`

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func (userHandler *UserHandler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	urlPath := r.URL.Path
	parts := strings.Split(urlPath, "/")
	idStr := parts[len(parts)-1]

	id, _ := strconv.Atoi(idStr)

	result, err := userHandler.s.GetUserBalance(id)
	if err != nil {
		http.Error(w, "Cannot find user balance", http.StatusBadRequest)
		return
	}

	s := fmt.Sprintf("%f", result)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(s))
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func (userHandler *UserHandler) TopUpBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		ID     int     `json:"id"`
		Amount float64 `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = userHandler.s.UpdateBalance(data.ID, data.Amount)
	if err != nil {
		http.Error(w, "Cannot update balance", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
