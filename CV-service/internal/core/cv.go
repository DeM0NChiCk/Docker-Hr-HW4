package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type CV struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName  string             `bson:"firstName,omitempty" json:"firstName"`
	LastName   string             `bson:"lastName,omitempty" json:"lastName"`
	City       string             `bson:"city,omitempty" json:"city"`
	Experience int32              `bson:"experience,omitempty" json:"experience"`
	Skills     []string           `bson:"skills,omitempty" json:"skills"`
}
