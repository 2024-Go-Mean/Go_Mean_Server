package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Worries struct {
	Id         primitive.ObjectID
	Title      string
	Content    string
	Nickname   string
	AiAdviceId Advices
	CategoryId Categories
	CreatedAt  primitive.DateTime
}

func getCurrentTime() time.Time {
	return time.Now()
}

func NewWorry(id primitive.ObjectID, title string, content string, nickname string, aiAdviceId Advices, categoryId Categories) *Worries {

	return &Worries{
		Id:         id,
		Title:      title,
		Content:    content,
		Nickname:   nickname,
		AiAdviceId: aiAdviceId,
		CategoryId: categoryId,
		CreatedAt:  getCurrentTime(),
	}
}

func (model *Worries) CollectionName() string {
	return "worries"
}
