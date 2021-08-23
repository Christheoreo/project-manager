package interfaces

import "github.com/Christheoreo/project-manager/models"

type ITagsRepository interface {
	Insert(tag models.NewTag, userId int) (id int, err error)
	GetById(id int) (tag models.Tag, err error)
	GetAllForUser(userID int) (tags []models.Tag, err error)
	Exists(name string, userID int) (exists bool, err error)
}

type ITagsService interface {
	Create(newTag models.NewTag, user models.User) (tag models.Tag, err error)
	GetAll(user models.User) (tags []models.Tag, err error)
	Get(id int) (tag models.Tag, err error)
}
