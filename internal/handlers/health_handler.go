package handlers

import (
	"log"
	"net/http"
)

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Сервер запущен успешно!")
}
