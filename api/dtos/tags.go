package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	NewTagDto struct {
		Name string `bson:"name,omitempty" json:"name"`
	}

	TagDto struct {
		ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		Name string             `bson:"name,omitempty" json:"name"`
	}
)
