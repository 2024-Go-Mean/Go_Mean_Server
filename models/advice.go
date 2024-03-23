package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Advices struct {
	Id         primitive.ObjectID
	CategoryId primitive.ObjectID
	AiAdvice   string
}

func NewAdvice(id primitive.ObjectID, categoryId primitive.ObjectID, aiAdvice string) *Advices {

	return &Advices{
		Id:         id,
		CategoryId: categoryId,
		AiAdvice:   aiAdvice,
	}
}

func (model *Advices) CollectionName() string {
	return "advices"
}
