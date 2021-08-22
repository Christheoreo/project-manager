package interfaces

import "github.com/Christheoreo/project-manager/dtos"

type ITagsRepository interface {
	Insert(tag dtos.NewTagDto, userId int) (id int, err error)
	GetById(id int) (tag dtos.TagDto, err error)
	GetAllForUser(userID int) (tags []dtos.TagDto, err error)
	Exists(name string, userID int) (exists bool, err error)
}

type ITagsService interface {
	Create(newTag dtos.NewTagDto, user dtos.UserDto) (tag dtos.TagDto, err error)
	GetAll(user dtos.UserDto) (tags []dtos.TagDto, err error)
	Get(id int) (tag dtos.TagDto, err error)
}
