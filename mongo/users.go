package mongo

import (
	"fmt"
	models "photon-backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(u *models.User) models.User {
	inserted, err := usersCollection.InsertOne(ctx, *u)
	if err != nil {
		fmt.Println("Error inserting user", err)
	}
	u.ID = inserted.InsertedID.(primitive.ObjectID)

	return *u
}

func GetUserByID(id primitive.ObjectID) (models.User, error) {
	var u models.User
	err := usersCollection.FindOne(ctx, primitive.M{"oauth_uid": id}).Decode(&u)
	if err != nil {
		return models.User{}, fmt.Errorf("error fetching user by ID: %w", err)
	}
	return u, nil
}
