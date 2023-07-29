package main

import (
	"bank-app/src/app/controllers/user"
	"net/http"
)

func main() {
	http.Handle("/", new(user.UserHandler))

	http.ListenAndServe(":3333", nil)
}
