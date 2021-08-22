package services

import (
	"errors"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/interfaces"
)

type ProjectsService struct {
	ProjectsRepository   interfaces.IProjectsRepository
	TagsRepository       interfaces.ITagsRepository
	PrioritiesRepository interfaces.IPrioritiesRepository
}

type (
	ProjectComponentToInsert struct {
		ID          int
		Title       string      `json:"title"`
		Description string      `json:"description"`
		Data        interface{} `json:"data"`
	}

	ProjectToInsert struct {
		ID          int
		UserID      int                        `json:"userId"`
		Title       string                     `json:"title"`
		Description string                     `json:"description"`
		Tags        []string                   `json:"tags"`
		Priority    string                     `json:"priority"`
		Components  []ProjectComponentToInsert `json:"components"`
	}
)

func (s *ProjectsService) Get(ID int, user dtos.UserDto) (dtos.ProjectDto, error) {
	project, err := s.ProjectsRepository.GetByID(ID)
	if err != nil {
		return project, errors.New("invalid project")
	}

	userID, err := s.ProjectsRepository.GetOwnerID(project.ID)
	if err != nil {
		return project, errors.New("invalid project")
	}

	if userID != user.ID {
		return project, errors.New("you do not have access to this project")
	}

	return project, nil
}

func (s *ProjectsService) All(user dtos.UserDto) ([]dtos.ProjectDto, error) {
	//
	return s.ProjectsRepository.GetByUser(user)
}

func (s *ProjectsService) Create(newProjectDto dtos.NewProjectDto, user dtos.UserDto) (project dtos.ProjectDto, err error) {
	id, err := s.ProjectsRepository.Insert(newProjectDto, user)
	if err != nil {
		return
	}
	return s.ProjectsRepository.GetByID(id)
}

func (s *ProjectsService) TitleTaken(title string, userID int) (bool, error) {
	return s.ProjectsRepository.TitleTaken(title, userID)

}
