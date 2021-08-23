package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Christheoreo/project-manager/interfaces"
	"github.com/Christheoreo/project-manager/models"
)

type (
	ProjectsHandler struct {
		ProjectsService   interfaces.IProjectsService
		TagsService       interfaces.ITagsService
		PrioritiesService interfaces.IPrioritiesService
	}
)

func (h *ProjectsHandler) validateNewProjectDto(newProject models.NewProjectDto, user models.UserDto) ([]string, error) {

	errorMessages := make([]string, 0)

	if len(newProject.Title) < 3 {
		e := "'title' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}

	if len(newProject.Description) < 3 {
		e := "'description' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}

	tags, err := h.TagsService.GetAll(user)

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

	priorityExists, err := h.PrioritiesService.Exists(newProject.PriorityID)

	if err != nil || !priorityExists {
		errorMessages = append(errorMessages, "PriorityID is invalid.")
	}

	// check to see if project name already exists.

	isTaken, err := h.ProjectsService.TitleTaken(newProject.Title, user.ID)

	if err != nil {
		errorMessages = append(errorMessages, "could not verify if project exists.")
	}

	if isTaken {
		errorMessages = append(errorMessages, "project with title already exists.")
	}
	if len(errorMessages) > 0 {
		err = errors.New("invalid DTO")
	}
	return errorMessages, err
}

func (p *ProjectsHandler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	user := getUser(r)
	var newProject models.NewProjectDto

	err := json.NewDecoder(r.Body).Decode(&newProject)
	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	errMessages, err := p.validateNewProjectDto(newProject, user)

	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, errMessages)
		return
	}

	project, err := p.ProjectsService.Create(newProject, user)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to create project."})
		return
	}

	returnObjectResponse(w, http.StatusCreated, project)
}

func (p *ProjectsHandler) GetMyProjects(w http.ResponseWriter, r *http.Request) {
	user := getUser(r)

	projects, err := p.ProjectsService.All(user)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to fetch projects."})
		return
	}
	returnObjectResponse(w, http.StatusCreated, projects)
}
