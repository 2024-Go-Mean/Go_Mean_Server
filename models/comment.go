package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment 모델 정의
type Comment struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	WorryID  string             `json:"worry_id,omitempty" bson:"worry_id,omitempty"`
	Nickname string             `json:"nickname,omitempty"`
	Comment  string             `json:"comment"`
}
