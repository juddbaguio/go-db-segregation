package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"title,omitempty"`
	LastName  string             `json:"lastName" bson:"author,omitempty"`
	Age       int                `json:"age" bson:"tags,omitempty"`
}
