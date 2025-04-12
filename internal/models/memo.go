package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Memo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	UserType  Role               `json:"userType" bson:"user_type"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
}
