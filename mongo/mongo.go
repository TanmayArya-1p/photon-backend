package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var usersCollection *mongo.Collection
var sessionsCollection *mongo.Collection
var ctx = context.Background()

func Connect() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING"))
	err := error(nil)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB", err)
	}
	usersCollection = client.Database("photon").Collection("users")
	sessionsCollection = client.Database("photon").Collection("sessions")

}

func ObjIDfromString(id string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error converting string to ObjectID", err)
	}
	return objID
}
