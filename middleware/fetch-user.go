package middleware

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	models "photon-backend/models"
	mongo "photon-backend/mongo"
	pebble "photon-backend/pebble"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	PASSWORD_KEY = os.Getenv("PASSWORD_KEY")
)

func CreateNewUserMapping(JWTdecoded models.UnpackedAccessToken) models.User {
	uid := JWTdecoded.UID
	acr := JWTdecoded.Acr

	password := uid + ":" + PASSWORD_KEY
	nh := sha256.New()
	nh.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", nh.Sum(nil))
	registrationResponse, err := pebble.Register(uid, hashedPassword[:10])
	if err != nil {
		fmt.Println("Error Registering User", err)
	}
	fmt.Println("Registration Response", registrationResponse)

	nU := models.User{
		ID:             mongo.ObjIDfromString(uid),
		Oauth_provider: acr,
		Created_at:     time.Now(),
		IsAlive:        false,
		InSession:      primitive.ObjectID{},
		Friends:        []primitive.ObjectID{},
		PebbleUID:      registrationResponse.UID,
		PebbleSecret:   registrationResponse.ClientSecret,
		PebblePassword: hashedPassword[:10],
		Email:          JWTdecoded.Email,
	}
	mongo.InsertUser(&nU)
	return nU
}

func FetchUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		unpackedAt := r.Context().Value(models.UnpackedAT)
		uid := unpackedAt.(models.UnpackedAccessToken).UID
		user, err := mongo.GetUserByID(mongo.ObjIDfromString(uid))
		if err != nil {
			fmt.Println("Need to Create a New User")
			CreateNewUserMapping(unpackedAt.(models.UnpackedAccessToken))
			user, err = mongo.GetUserByID(mongo.ObjIDfromString(uid))
			if err != nil {
				http.Error(w, "error fetching user after creating", http.StatusInternalServerError)
			}
			r = r.WithContext(context.WithValue(r.Context(), models.UserContextKey, user))
		} else {
			r = r.WithContext(context.WithValue(r.Context(), models.UserContextKey, user))
		}
		next.ServeHTTP(w, r)
	})
}
