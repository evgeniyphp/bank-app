package user

import "net/http"

type UserHandler struct{}

func (userHandler *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User handler"))
}
