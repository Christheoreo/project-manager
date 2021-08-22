package services

import (
	"errors"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/interfaces"
)

type TagsService struct {
	TagsRepository interfaces.ITagsRepository
}

func (s *TagsService) Create(newTag dtos.NewTagDto, user dtos.UserDto) (tag dtos.TagDto, err error) {

	taken, _ := s.TagsRepository.Exists(newTag.Name, user.ID)

	if taken {
		err = errors.New("tag name already exists")
		return
	}

	insertedId, err := s.TagsRepository.Insert(newTag, user.ID)

	if err != nil {
		err = errors.New("unable to create tag")
		return
	}
	tag, _ = s.TagsRepository.GetById(insertedId)

	return
}
func (s *TagsService) Get(id int) (tag dtos.TagDto, err error) {
	return s.TagsRepository.GetById(id)
}
func (s *TagsService) GetAll(user dtos.UserDto) (tags []dtos.TagDto, err error) {
	return s.TagsRepository.GetAllForUser(user.ID)
}
