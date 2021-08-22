package services

import (
	"errors"
	"fmt"

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

func (s *ProjectsService) Get(ID int) (dtos.ProjectDto, error) {
	// @todo
	// need to check if A) ID exists and B) User owns the project.
	return s.ProjectsRepository.GetByID(ID)
}

func (s *ProjectsService) All(user dtos.UserDto) ([]dtos.ProjectDto, error) {
	//
	return s.ProjectsRepository.GetByUser(user)
}
func (s *ProjectsService) ValidateNewProjectDto(newProject dtos.NewProjectDto, user dtos.UserDto) ([]string, error) {

	errorMessages := make([]string, 0)

	if len(newProject.Title) < 3 {
		e := "'title' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}

	if len(newProject.Description) < 3 {
		e := "'description' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}

	tags, err := s.TagsRepository.GetAllForUser(user.ID)

	if err != nil {
		errorMessages = append(errorMessages, "invalid tags")
	}

	for _, tagID := range newProject.TagIDs {
		exists := false
		for _, tag := range tags {
			if tag.ID == tagID {
				exists = true
				break
			}
		}
		if !exists {
			e := fmt.Sprintf("'%d' is not a recognised tag", tagID)
			errorMessages = append(errorMessages, e)
		}
	}

	// continue validation stuff here

	priorityExists, err := s.PrioritiesRepository.Exists(newProject.PriorityID)

	if err != nil {
		errorMessages = append(errorMessages, "PriorityID is bad.")
	} else if !priorityExists {
		errorMessages = append(errorMessages, "PriorityID is invalid.")
	}

	// check to see if project name already exists.

	isTaken, err := s.ProjectsRepository.TitleTaken(newProject.Title, user.ID)

	if err != nil {
		errorMessages = append(errorMessages, "could not verify if project exists.")
	}

	if isTaken {
		errorMessages = append(errorMessages, "project name already exists.")
	}
	if len(errorMessages) > 0 {
		err = errors.New("invalid DTO")
	}
	return errorMessages, err
}

func (s *ProjectsService) Create(newProjectDto dtos.NewProjectDto, user dtos.UserDto) (project dtos.ProjectDto, err error) {
	id, err := s.ProjectsRepository.Insert(newProjectDto, user)
	if err != nil {
		return
	}
	return s.ProjectsRepository.GetByID(id)
}
