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
		err = errors.New("email is taken")
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
func (s *TagsService) ValidateNewTagDto(newTag dtos.NewTagDto) ([]string, error) {

	errorMessages := make([]string, 0)
	var err error
	if len(newTag.Name) < 3 {
		e := "'name' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}

	if len(errorMessages) > 0 {
		err = errors.New("invalid DTO")
	}
	return errorMessages, err
}
