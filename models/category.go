package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories struct {
	ID       primitive.ObjectID
	Category string
}
