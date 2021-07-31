package models

import (
	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Collection *mongo.Collection
}

type UserToInsert struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName,omitempty" json:"firstName"`
	LastName  string             `bson:"lastName,omitempty" json:"lastName"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password"`
}

func (u *User) HasEmailBeenTaken(email string) (bool, error) {
	count, err := u.Collection.CountDocuments(getContext(), bson.M{"email": email})

	if err != nil {
		return false, err
	}
	return count > 0, err

}

func (u *User) Insert(user dtos.NewUserDto) (id primitive.ObjectID, err error) {

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return
	}

	userToInsert := UserToInsert{
		ID:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  passwordHash, // needs encrypting
	}
	res, err := u.Collection.InsertOne(getContext(), userToInsert)

	if err != nil {
		return
	}

	id = res.InsertedID.(primitive.ObjectID)
	return
}

func (u *User) GetById(userId primitive.ObjectID) (user dtos.UserDto, err error) {
	err = u.Collection.FindOne(getContext(), bson.M{"_id": userId}).Decode(&user)
	return
}

func (u *User) GetByEmail(email string) (user dtos.UserDto, err error) {
	err = u.Collection.FindOne(getContext(), bson.M{"email": email}).Decode(&user)
	return
}

func (u *User) ValidateUserCredentials(authLogin dtos.AuthLoginDto) (valid bool, err error) {

	var fullUser UserToInsert
	err = u.Collection.FindOne(getContext(), bson.M{"email": authLogin.Email}).Decode(&fullUser)

	if err != nil {
		return
	}
	valid = utils.CheckPasswordHash(authLogin.Password, fullUser.Password)
	return
}

func (u *User) GetIDAsString(userId primitive.ObjectID) string {
	return userId.Hex()
}
