package dtos

type (
	NewProjectDto struct {
		Title        string   `bson:"title,omitempty" json:"title"`
		Description  string   `bson:"description,omitempty" json:"description"`
		TypeStringID string   `bson:"typeId,omitempty" json:"typeId"`
		TagStringIDs []string `bson:"tagIds,omitempty" json:"tagIds"`
	}
)
