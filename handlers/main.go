package handlers

import (
	"encoding/json"
	"net/http"
	"photon-backend/models"
	"photon-backend/mongo"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	user := r.Context().Value(models.UserContextKey).(models.User)
	if user.InSession != mongo.ObjIDfromString("000000000000000000000000") && !user.IsAlive {
		user.IsAlive = true
		user, _ = mongo.UpdateUser(user)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
