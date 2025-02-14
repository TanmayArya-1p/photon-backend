package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendInvite struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	To        primitive.ObjectID `bson:"to" json:"to"`
	From      primitive.ObjectID `bson:"from" json:"from"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type FriendInviteResponse struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	InviteID primitive.ObjectID `bson:"invite_id" json:"invite_id"`
	Status   string             `bson:"status" json:"status"`
}
