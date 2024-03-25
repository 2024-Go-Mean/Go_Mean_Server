package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	WorryID  int                `json:"worry_id,omitempty" bson:"worry_id,omitempty"`
	Comment  string             `json:"comment,omitempty" bson:"comment,omitempty"`
	Nickname string             `json:"nickname,omitempty" bson:"nickname,omitempty"`
}
