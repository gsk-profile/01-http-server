package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gsklearn2025/go/01-http-server/internal/db"
	"github.com/gsklearn2025/go/01-http-server/internal/model"
)

var demoUser = model.User{
	Email:    "demo@example.com",
	Password: "pass123",
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds model.User
	err := json.NewDecoder(r.Body).Decode((&creds))
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var storedPassword string
	err = db.DB.QueryRow("SELECT password FROM users WHERE email = $1", creds.Email).Scan(&storedPassword)

	if err == sql.ErrNoRows {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if creds.Password == storedPassword {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login Successful"))

	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}

}
