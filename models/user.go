package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID   `bson:"_id" json:"id"`
	Email          string               `bson:"email" json:"email"`
	Oauth_uid      string               `bson:"oauth_uid" json:"oauth_uid"`
	Oauth_provider string               `bson:"oauth_provider" json:"oauth_provider"`
	Created_at     time.Time            `bson:"created_at" json:"created_at"`
	IsAlive        bool                 `bson:"is_alive" json:"is_alive"`
	InSession      primitive.ObjectID   `bson:"in_session" json:"in_session"`
	Friends        []primitive.ObjectID `bson:"friends" json:"friends"`
	PebbleUID      string               `bson:"pebble_uid" json:"pebble_uid"`
	PebbleSecret   string               `bson:"pebble_secret" json:"pebble_secret"`
}
