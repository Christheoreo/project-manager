package interfaces

import "github.com/Christheoreo/project-manager/models"

type ITagsRepository interface {
	Insert(tag models.NewTagDto, userId int) (id int, err error)
	GetById(id int) (tag models.TagDto, err error)
	GetAllForUser(userID int) (tags []models.TagDto, err error)
	Exists(name string, userID int) (exists bool, err error)
}

type ITagsService interface {
	Create(newTag models.NewTagDto, user models.UserDto) (tag models.TagDto, err error)
	GetAll(user models.UserDto) (tags []models.TagDto, err error)
	Get(id int) (tag models.TagDto, err error)
}
