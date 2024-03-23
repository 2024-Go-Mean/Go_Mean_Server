package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Advices struct {
	Id         primitive.ObjectID
	CategoryId primitive.ObjectID
	AiAdvice   string
}

type Categories struct {
	Id       primitive.ObjectID
	Category string
}
