package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/interfaces"
	"github.com/Christheoreo/project-manager/middleware"
)

type (
	ProjectsHandler struct {
		ProjectsService   interfaces.IProjectsService
		TagsService       interfaces.ITagsService
		PrioritiesService interfaces.IPrioritiesService
	}
)

func (p *ProjectsHandler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)
	var newProject dtos.NewProjectDto

	err := json.NewDecoder(r.Body).Decode(&newProject)
	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	errMessages, err := p.ProjectsService.ValidateNewProjectDto(newProject, user)

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
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)

	projects, err := p.ProjectsService.All(user)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to fetch projects."})
		return
	}
	returnObjectResponse(w, http.StatusCreated, projects)
}
