package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	NewUserDto struct {
		FirstName       string `bson:"firstName,omitempty" json:"firstName"`
		LastName        string `bson:"lastName,omitempty" json:"lastName"`
		Email           string `bson:"email,omitempty" json:"email"`
		Password        string `bson:"password,omitempty" json:"password"`
		PasswordConfirm string `json:"passwordConfirm"`
	}

	UserDto struct {
		ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		FirstName string             `bson:"firstName,omitempty" json:"firstName"`
		LastName  string             `bson:"lastName,omitempty" json:"lastName"`
		Email     string             `bson:"email,omitempty" json:"email"`
	}
)
