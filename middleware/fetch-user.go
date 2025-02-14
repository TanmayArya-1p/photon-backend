package middleware

import (
	"fmt"
	"net/http"
)

func FetchUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value("uid")
		fmt.Println("REQUEST SENT BY ", uid)
		next.ServeHTTP(w, r)
	})
}
