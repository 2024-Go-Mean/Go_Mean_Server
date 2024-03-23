package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories struct {
	Id       primitive.ObjectID
	Category string
}
