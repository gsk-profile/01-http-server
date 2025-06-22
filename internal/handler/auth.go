package handler

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var demoUser = User{
	Email:    "demo@example.com",
	Password: "pass123",
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds User
	err := json.NewDecoder(r.Body).Decode((&creds))
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if creds.Email == demoUser.Email && creds.Password == demoUser.Password {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login Successful"))

	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}

}
