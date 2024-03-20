package models

import (
	"github.com/pelletier/go-toml/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Worries struct {
	ID         primitive.ObjectID
	Title      string
	Content    string
	Nickname   string
	AiAdviceId primitive.ObjectID
	CategoryId primitive.ObjectID
	CreatedAt  toml.LocalDateTime
}

type Advices struct {
	ID         primitive.ObjectID
	CategoryId primitive.ObjectID
	AiAdvice   string
}

type Categories struct {
	ID       primitive.ObjectID
	Category string
}
