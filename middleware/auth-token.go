package middleware

import (
	"context"
	"net/http"
	auth "photon-backend/auth"
	"photon-backend/models"
)

func AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header missing", http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len("Bearer "):]
		valid, err, unpackedjwt := auth.ValidateAuthToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), models.UnpackedAT, unpackedjwt))
		next.ServeHTTP(w, r)
	})
}
