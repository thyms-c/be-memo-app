package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Counter struct {
	ID    bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	Value int           `json:"value" bson:"value"`
}
