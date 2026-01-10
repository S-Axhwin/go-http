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
	fmt.Fprintln(w, "this is going so be home page")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invaild Inputs", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, user := range users {
		if user.Username == newUser.Username {
			http.Error(w, "User Already Exists", http.StatusConflict)
			return
		}
	}
	users = append(users, newUser)

	fmt.Fprintln(w, "user created done: ", newUser.Username)
}

func main() {
	http.HandleFunc("GET /", HomeHandler)
	http.HandleFunc("GET /users", GetUsers)
	http.HandleFunc("POST /users", InsertUser)
	http.ListenAndServe(":8080", nil)

}
