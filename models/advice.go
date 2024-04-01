package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Advices struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	CategoryId primitive.ObjectID `json:"category_id"`
	AiAdvice   string             `json:"ai_advice"`
}
