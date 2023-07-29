package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.Handle("/", helloWorld)

	http.ListenAndServe(":3333", nil)
}
