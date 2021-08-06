package models

import (
	"github.com/Christheoreo/project-manager/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Tag struct {
		Collection *mongo.Collection
	}

	TagToInsert struct {
		ID     primitive.ObjectID `bson:"_id,omitempty"`
		Name   string             `bson:"name,omitempty" json:"name"`
		UserID primitive.ObjectID `bson:"userId,omitempty"`
	}
)

func (t *Tag) GetById(tagId primitive.ObjectID) (tag dtos.TagDto, err error) {
	err = t.Collection.FindOne(getContext(), bson.M{"_id": tagId}).Decode(&tag)
	return
}

func (t *Tag) HasTagBeenTakenByUser(name string, userId primitive.ObjectID) (bool, error) {
	count, err := t.Collection.CountDocuments(getContext(), bson.M{"name": name, "userId": userId})

	if err != nil {
		return false, err
	}
	return count > 0, err

}

func (t *Tag) Insert(tag dtos.NewTagDto, user dtos.UserDto) (id primitive.ObjectID, err error) {

	tagToInsert := TagToInsert{
		ID:     primitive.NewObjectID(),
		Name:   tag.Name,
		UserID: user.ID,
	}
	res, err := t.Collection.InsertOne(getContext(), tagToInsert)

	if err != nil {
		return
	}

	id = res.InsertedID.(primitive.ObjectID)
	return
}
