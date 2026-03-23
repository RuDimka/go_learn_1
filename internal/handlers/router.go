package handlers

import (
	"net/http"
	"test/app/internal/service"
)

type Handler struct {
	UserService *service.UserService
}

func NewRouter(userService *service.UserService) *http.ServeMux {
	h := &Handler{UserService: userService}

	mux := http.NewServeMux()
	mux.HandleFunc("/user", h.HelloHandler)
	mux.HandleFunc("/health", h.HealthHandler)
	mux.HandleFunc("/time", h.TimeHandler)
	mux.HandleFunc("/users", h.GetUsers)
	mux.HandleFunc("/users/create", h.CreateUser)
	return mux
}
