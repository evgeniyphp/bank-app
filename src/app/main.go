package main

import (
	sqlite3 "bank-app/pkg/database/sqlite3"
	userHandler "bank-app/src/app/controllers/user"
	"bank-app/src/app/controllers/user/services"
	"net/http"

	userModel "bank-app/src/app/controllers/user/models"

	"os"
)

func main() {
	storage, err := sqlite3.New("./storage")
	if err != nil {
		os.Exit(1)
	}
	userRepository := userModel.NewUserRepository(storage.DB)
	userService := services.NewUserService(userRepository)
	uH := userHandler.NewUserHandler(userService)
	http.HandleFunc("/users", uH.CreateUser)
	http.HandleFunc("/users", uH.GetUserBalance)
	//	http.HandleFunc("/",)

	http.ListenAndServe(":3333", nil)
}
