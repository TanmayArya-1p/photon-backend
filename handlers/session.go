package handlers

import (
	"encoding/json"
	"net/http"
	models "photon-backend/models"
	mng "photon-backend/mongo"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateSessionEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	user := r.Context().Value(models.UserContextKey).(models.User)

	var req models.UserSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	session := models.Session{
		ID:         mng.ObjIDfromString(req.SessionID),
		SessionKey: req.SessionKey,
		Users:      []primitive.ObjectID{user.ID},
		IsAlive:    true,
		Created_at: time.Now(),
	}

	newses := mng.CreateSession(&session)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newses); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func JoinSessionEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	user := r.Context().Value(models.UserContextKey).(models.User)

	var req models.UserSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	session := mng.GetSessionByID(mng.ObjIDfromString(req.SessionID))

	if session.SessionKey != req.SessionKey {
		http.Error(w, "Invalid session key", http.StatusUnauthorized)
		return
	}
	session.Users = append(session.Users, user.ID)
	session = mng.UpdateSession(session)
	err := json.NewEncoder(w).Encode(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
