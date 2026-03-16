package handlers

import (
	"net/http"
	"time"
)

func (h *Handler) TimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("15:04:05 02.01.2006")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Тукущее время: " + currentTime))
}
