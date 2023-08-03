package main

import (
	"bank-app/pkg/database/sqlite3"
	"bank-app/src/app/controllers/purchase"
	purchaseModel "bank-app/src/app/controllers/purchase/models"
	purchaseService "bank-app/src/app/controllers/purchase/services"
	transactionModel "bank-app/src/app/controllers/transcation/models"
	transactionService "bank-app/src/app/controllers/transcation/services"
	userHandler "bank-app/src/app/controllers/user"
	userModel "bank-app/src/app/controllers/user/models"
	userService "bank-app/src/app/controllers/user/services"
	"net/http"

	"os"
)

func handleEndpoint(handler *userHandler.UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handler.CreateUser(w, r)
		} else if r.Method == "GET" {
			handler.GetUserBalance(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func main() {
	storage, err := sqlite3.New("./storage")
	if err != nil {
		os.Exit(1)
	}
	userRepository := userModel.New(storage.DB)
	goodRepository := purchaseModel.New(storage.DB)
	transactionRepository := transactionModel.New(storage.DB)

	tS := transactionService.New(transactionRepository)
	uS := userService.New(userRepository, tS)
	pS := purchaseService.New(goodRepository, userRepository, transactionRepository)

	uH := userHandler.NewUserHandler(uS)
	sH := purchase.New(pS)

	mux := http.NewServeMux()

	mux.HandleFunc("/users", handleEndpoint(uH))
	mux.HandleFunc("/update-balance", uH.TopUpBalance)
	mux.HandleFunc("/buy-good", sH.BuyGood)
	mux.HandleFunc("/good", sH.CreateGood)

	http.ListenAndServe(":3333", mux)
}
