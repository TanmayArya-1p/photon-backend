package handlers

import (
	"encoding/json"
	"net/http"
	"photon-backend/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	user := r.Context().Value(models.UserContextKey).(models.User)
	w.Header().Set("Content-Type", "application/json")
	user.PebblePassword = ""
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
