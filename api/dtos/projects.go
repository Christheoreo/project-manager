package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	NewProjectComponentDto struct {
		Title       string      `bson:"title,omitempty" json:"title"`
		Description string      `bson:"description,omitempty" json:"description"`
		Data        interface{} `bson:"data,omitempty" json:"data"`
	}

	ProjectComponentDto struct {
		ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		Title       string             `bson:"title,omitempty" json:"title"`
		Description string             `bson:"description,omitempty" json:"description"`
		Data        interface{}        `bson:"data,omitempty" json:"data"`
	}

	NewProjectDto struct {
		Title       string                   `bson:"title,omitempty" json:"title"`
		Description string                   `bson:"description,omitempty" json:"description"`
		Tags        []string                 `bson:"tags,omitempty" json:"tags"`
		Priority    string                   `bson:"priority,omitempty" json:"priority"`
		Components  []NewProjectComponentDto `bson:"components,omitempty" json:"components"`
	}

	ProjectDto struct {
		ID          primitive.ObjectID    `bson:"_id,omitempty" json:"id"`
		Title       string                `bson:"title,omitempty" json:"title"`
		Description string                `bson:"description,omitempty" json:"description"`
		Tags        []string              `bson:"tags,omitempty" json:"tags"`
		Priority    string                `bson:"priority,omitempty" json:"priority"`
		Components  []ProjectComponentDto `bson:"components,omitempty" json:"components"`
	}
)
