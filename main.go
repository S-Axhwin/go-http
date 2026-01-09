package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users []User
var mu sync.Mutex

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO:  Impliment Home page
	fmt.Fprintln(w, "this is going to be home page")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Fprintln(w, "get the users in this as a object")

}

func InsertUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invaild Inputs", http.StatusBadRequest)
		return
	}

	mu.Lock()
	users = append(users, newUser)
	mu.Unlock()

	fmt.Fprintln(w, "user created done: ", newUser.Username)
}

func main() {
	http.HandleFunc("GET /", HomeHandler)
	http.HandleFunc("GET /users", GetUsers)
	http.HandleFunc("POST /users", InsertUser)
	http.ListenAndServe(":8080", nil)

}
