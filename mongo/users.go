package mongo

import (
	"fmt"
	models "photon-backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(u *models.User) (models.User, error) {
	inserted, err := usersCollection.InsertOne(ctx, *u)
	if err != nil {
		fmt.Println("Error inserting user", err)
		return models.User{}, fmt.Errorf("error inserting user: %w", err)
	}
	u.ID = inserted.InsertedID.(primitive.ObjectID)

	return *u, nil
}

func GetUserByID(id primitive.ObjectID) (models.User, error) {
	var u models.User
	err := usersCollection.FindOne(ctx, primitive.M{"_id": id}).Decode(&u)
	if err != nil {
		return models.User{}, fmt.Errorf("error fetching user by ID: %w", err)
	}
	return u, nil
}
