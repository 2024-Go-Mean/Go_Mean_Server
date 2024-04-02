package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Category string             `json:"category"`
}
