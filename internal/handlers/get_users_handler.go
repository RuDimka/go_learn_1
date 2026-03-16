package handlers

import (
	"encoding/json"
	"net/http"
	"test/app/internal/models"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Не верный метод", http.StatusMethodNotAllowed)
		return
	}
	var users []models.User

	if err := h.db.Order("created_at desc").Find(&users).Error; err != nil {
		http.Error(w, "Ошибка при получении всех пользователей : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
