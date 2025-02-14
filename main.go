package main

import (
	"fmt"
	"net/http"
	"os"
	"photon-backend/handlers"
	middleware "photon-backend/middleware"
	mongo "photon-backend/mongo"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mongo.Connect()
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.Handle("/login", middleware.AuthTokenMiddleware(middleware.FetchUserMiddleware(http.HandlerFunc(handlers.Login))))

	fmt.Println("Starting server at port", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
