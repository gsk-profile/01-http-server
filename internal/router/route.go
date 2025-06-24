package router

import (
	"net/http"

	"github.com/gsklearn2025/go/01-http-server/internal/handler"
)

func RegisterRoutes() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/auth/login", handler.LoginHandler)
	http.HandleFunc("/auth/logout", handler.LogoutHandler)
}
