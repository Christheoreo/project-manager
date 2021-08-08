package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/middleware"
	"github.com/Christheoreo/project-manager/models"
)

type (
	ProjectsHandler struct {
		ProjectModel models.Project
		TagModel     models.Tag
	}
)

/**
Loop through and check if the components are valid.
*/
func (p *ProjectsHandler) validateNewProjectDto(body io.Reader, user dtos.UserDto) (newProject dtos.NewProjectDto, errorMessages []string, err error) {

	err = json.NewDecoder(body).Decode(&newProject)
	if err != nil {
		return
	}

	errorMessages = make([]string, 0)

	if len(newProject.Title) < 3 {
		e := "'title' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}

	if len(newProject.Description) < 3 {
		e := "'description' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}

	for _, tagID := range newProject.TagIDs {
		exists, err := p.TagModel.DoesTagExistForUserByID(tagID, user.ID)

		if err != nil {
			errorMessages = append(errorMessages, "TagID is bad.")
			break
		}

		if !exists {
			e := fmt.Sprintf("'%d' is not a recognised tag", tagID)
			errorMessages = append(errorMessages, e)
		}
		// contiue validation stuff here
		// then test the create endpoints and etc
	}

	if len(errorMessages) > 0 {
		err = errors.New("invalid DTO")
	}
	return
}

func (p *ProjectsHandler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)
	newProject, errMessages, err := p.validateNewProjectDto(r.Body, user)

	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, errMessages)
		return
	}

	// Check if a user with that email exists

	taken, err := p.ProjectModel.HasProjectBeenTakenByUser(newProject.Title, user.ID)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Could not verify if project already exists."})
		return
	}

	if taken {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"You already have a project with that title."})
		return
	}

	insertedId, err := p.ProjectModel.Insert(newProject, user)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to create project."})
		return
	}

	project, err := p.ProjectModel.GetById(insertedId)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to fetch created project."})
		return
	}
	returnObjectResponse(w, http.StatusCreated, project)
}

func (p *ProjectsHandler) GetMyProjects(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)

	projects, err := p.ProjectModel.GetByUser(user)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to fetch projects."})
		return
	}
	returnObjectResponse(w, http.StatusCreated, projects)
}
