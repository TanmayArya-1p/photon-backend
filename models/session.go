package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID         primitive.ObjectID   `bson:"_id" json:"id"`
	SessionKey string               `bson:"session_key" json:"session_key"`
	Users      []primitive.ObjectID `bson:"users" json:"users"`
	IsAlive    bool                 `bson:"is_alive" json:"is_alive"`
	Created_at time.Time            `bson:"created_at" json:"created_at"`
}
