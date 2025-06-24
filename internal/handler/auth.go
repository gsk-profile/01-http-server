package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gsklearn2025/go/01-http-server/internal/db"
	"github.com/gsklearn2025/go/01-http-server/internal/model"
)

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
		token := uuid.New().String()
		sessionId := uuid.New().String()
		now := time.Now()
		expiresAt := now.Add(24 * time.Hour)

		_, err = db.DB.Exec(
			"INSERT INTO SESSIONS (ID, EMAIL, TOKEN, CREATED_AT, EXPIRES_AT) VALUES ($1, $2, $3, $4, $5)",
			sessionId, creds.Email, token, now, expiresAt,
		)

		if err != nil {
			http.Error(w, "Failed to create session", http.StatusInternalServerError)
			return
		}

		_, err = db.DB.Exec(
			"INSERT INTO AUDIT_LOG (EMAIL, ACTION, TIMESTAMP, DETAILS) VALUES ($1, $2, $3, $4)",
			creds.Email, "LOGIN", now, "User logged in successfully",
		)

		if err != nil {
			http.Error(w, "Failed to create audit log", http.StatusInternalServerError)
			return
		}

		resp := map[string]string{"token": token}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Expect token in Authorization header: "Bearer <token>"
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) < 8 {
		http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
		return
	}
	token := authHeader[7:] // Remove "Bearer "

	// Fetch the session
	var email string
	err := db.DB.QueryRow(
		"SELECT email FROM sessions WHERE token = $1", token).Scan(&email)

	if err == sql.ErrNoRows {
		http.Error(w, "Unable to fetch session details", http.StatusInternalServerError)
		return
	}

	// Audit log
	db.DB.Exec(
		"INSERT INTO audit_log (email, action, timestamp, details) VALUES ($1, $2, $3, $4)",
		email, "LOGOUT", time.Now(), "User logged out",
	)

	// even if audit log fails, we should continue to delete the session

	// Delete session
	res, err := db.DB.Exec("DELETE FROM sessions WHERE token = $1", token)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
