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

func GetUserByID(id primitive.ObjectID) models.User {
	var u models.User
	err := usersCollection.FindOne(ctx, primitive.M{"_id": id}).Decode(&u)
	if err != nil {
		fmt.Println("Error fetching user by ID", err)
	}
	return u
}
