package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO:  Impliment Home page
	fmt.Fprintln(w, "Welcome to go lang")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: Impliment Global State for User
	// FIXME: This is a hack
	fmt.Fprintln(w, "get all users")

}

func main() {
	http.HandleFunc("GET /", HomeHandler)
	http.HandleFunc("GET /users", GetUsers)
	http.ListenAndServe(":8080", nil)
}
