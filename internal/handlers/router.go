package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewRouter(db *gorm.DB) *http.ServeMux {
	h := &Handler{db: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/user", h.HelloHandler)
	mux.HandleFunc("/health", h.HealthHandler)
	mux.HandleFunc("/time", h.TimeHandler)
	mux.HandleFunc("/users", h.GetUsers)
	return mux
}
