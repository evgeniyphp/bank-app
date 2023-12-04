package main

import (
	"bank-app/pkg/database/sqlite3"
	purchaseController "bank-app/src/app/controllers/purchase"
	purchaseModel "bank-app/src/app/controllers/purchase/models"
	purchaseService "bank-app/src/app/controllers/purchase/services"
	transactionModel "bank-app/src/app/controllers/transaction/models"
	transactionService "bank-app/src/app/controllers/transaction/services"
	userController "bank-app/src/app/controllers/user"
	userModel "bank-app/src/app/controllers/user/models"
	userService "bank-app/src/app/controllers/user/services"

	"net/http"

	"os"
)

func main() {
	storage, err := sqlite3.New("./storage")

	if err != nil {
		os.Exit(1)
	}

	userRepository := userModel.NewUserRepository(storage.DB)
	goodRepository := purchaseModel.NewGoodRepository(storage.DB)
	transactionRepository := transactionModel.NewTransactionRepository(storage.DB)

	tS := transactionService.NewTransactionService(transactionRepository)
	uS := userService.NewUserService(userRepository, tS)
	pS := purchaseService.NewPurchaseService(goodRepository, userRepository, transactionRepository)

	// handlers
	uH := userController.NewController(uS)
	sH := purchaseController.NewPurchaseController(pS)

	mux := http.NewServeMux()

	mux.HandleFunc("/users", uH.CreateUser)
	mux.HandleFunc("/users/", uH.GetUserBalance)
	mux.HandleFunc("/update-balance", uH.TopUpBalance)
	mux.HandleFunc("/buy-good", sH.BuyGood)
	mux.HandleFunc("/good", sH.CreateGood)

	http.ListenAndServe(":3333", mux)
}
