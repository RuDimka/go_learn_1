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

	err := h.UserService.GetUsers(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Список пользователей успешно получен",
		"users":   users,
	})
}
