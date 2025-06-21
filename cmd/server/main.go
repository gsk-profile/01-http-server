package main

import (
	"fmt"
	"net/http"

	"github.com/gsklearn2025/go/01-http-server/internal/handler"
)

func main() {
	fmt.Println("Hello Go")
	http.HandleFunc("/", handler.HomeHandler)
	http.ListenAndServe(":8080", nil)
}
