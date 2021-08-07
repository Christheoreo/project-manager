package models

import (
	"fmt"

	"github.com/Christheoreo/project-manager/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Project struct {
		Collection *mongo.Collection
	}

	ProjectComponentToInsert struct {
		ID          primitive.ObjectID `bson:"_id,omitempty"`
		Title       string             `bson:"title,omitempty" json:"title"`
		Description string             `bson:"description,omitempty" json:"description"`
		Data        interface{}        `bson:"data,omitempty" json:"data"`
	}

	ProjectToInsert struct {
		ID          primitive.ObjectID         `bson:"_id,omitempty"`
		UserID      primitive.ObjectID         `bson:"userId,omitempty" json:"userId"`
		Title       string                     `bson:"title,omitempty" json:"title"`
		Description string                     `bson:"description,omitempty" json:"description"`
		Tags        []string                   `bson:"tags,omitempty" json:"tags"`
		Priority    string                     `bson:"priority,omitempty" json:"priority"`
		Components  []ProjectComponentToInsert `bson:"components,omitempty" json:"components"`
	}
)

func (p *Project) GetById(tagId primitive.ObjectID) (project dtos.ProjectDto, err error) {
	err = p.Collection.FindOne(getContext(), bson.M{"_id": tagId}).Decode(&project)
	return
}
func (p *Project) GetByUser(user dtos.UserDto) (projects []dtos.ProjectDto, err error) {
	cursor, err := p.Collection.Find(getContext(), bson.M{"userId": user.ID})

	ctx := getContext()
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &projects)

	if err != nil {
		return
	}

	for i := 0; i < len(projects); i++ {
		var project dtos.ProjectDto = projects[i]

		components := project.Components

		newComponents := make([]dtos.ProjectComponentDto, len(project.Components))

		for _, component := range components {
			newComponents[i] = dtos.ProjectComponentDto{
				ID:          component.ID,
				Title:       component.Title,
				Description: component.Description,
				Data:        "null",
			}

		}
	}
	return
}

func (p *Project) HasProjectBeenTakenByUser(title string, userId primitive.ObjectID) (bool, error) {
	count, err := p.Collection.CountDocuments(getContext(), bson.M{"title": title, "userId": userId})

	if err != nil {
		return false, err
	}
	return count > 0, err

}

func (p *Project) Insert(project dtos.NewProjectDto, user dtos.UserDto) (id primitive.ObjectID, err error) {
	projectToInsert := &ProjectToInsert{
		ID:          primitive.NewObjectID(),
		UserID:      user.ID,
		Title:       project.Title,
		Description: project.Description,
		Tags:        project.Tags,
		Priority:    project.Priority,
		Components:  []ProjectComponentToInsert{},
	}
	// Need to add the ids to the component.
	for _, component := range project.Components {
		projectToInsert.Components = append(projectToInsert.Components, ProjectComponentToInsert{
			ID:          primitive.NewObjectID(),
			Title:       component.Title,
			Description: component.Description,
			Data:        component.Data,
		})
	}
	res, err := p.Collection.InsertOne(getContext(), projectToInsert)
	if err != nil {
		fmt.Println(err)
		return
	}

	id = res.InsertedID.(primitive.ObjectID)
	return
}
