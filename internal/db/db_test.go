package db

import (
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	// Set environment variables for testing (or ensure .env is present)
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "youruser")
	os.Setenv("DB_PASSWORD", "yourpassword")
	os.Setenv("DB_NAME", "yourdb")

	Connect()

	if DB == nil {
		t.Fatal("DB connection is nil")
	}

	err := DB.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
