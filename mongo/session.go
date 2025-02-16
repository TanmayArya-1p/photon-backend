package mongo

import (
	models "photon-backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateSession(s *models.Session) models.Session {
	inserted, err := sessionsCollection.InsertOne(ctx, *s)
	if err != nil {
		panic(err)
	}
	s.ID = inserted.InsertedID.(primitive.ObjectID)
	return *s
}

func GetSessionByID(id primitive.ObjectID) models.Session {
	var s models.Session
	err := sessionsCollection.FindOne(ctx, primitive.M{"_id": id}).Decode(&s)
	if err != nil {
		panic(err)
	}
	return s
}

func UpdateSession(s models.Session) models.Session {
	_, err := sessionsCollection.ReplaceOne(ctx, primitive.M{"_id": s.ID}, s)
	if err != nil {
		panic(err)
	}
	return s
}
