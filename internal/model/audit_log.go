package model

import "time"

type AuditLog struct {
	ID        int       `json:"id"`
	UserEmail string    `json:"user_email"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	Details   string    `json:"details"`
}
