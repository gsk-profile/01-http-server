package main

import (
	"log"
	"net/http"

	"github.com/gsklearn2025/go/01-http-server/internal/db"
	"github.com/gsklearn2025/go/01-http-server/internal/router"
)

func main() {
	db.Connect() // Initialize the database connection

	router.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
