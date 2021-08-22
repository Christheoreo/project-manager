package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/interfaces"
	"github.com/Christheoreo/project-manager/middleware"
)

type (
	TagsHandler struct {
		TagsService interfaces.ITagsService
	}
)

func (h *TagsHandler) validateNewTagDto(newTag dtos.NewTagDto) ([]string, error) {

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

func (h *TagsHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)
	var newTag dtos.NewTagDto

	err := json.NewDecoder(r.Body).Decode(&newTag)
	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	errMessages, err := h.validateNewTagDto(newTag)

	if err != nil {
		returnStandardResponse(w, http.StatusUnprocessableEntity, errMessages)
		return
	}

	tag, err := h.TagsService.Create(newTag, user)
	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	returnObjectResponse(w, http.StatusCreated, tag)
}

func (h *TagsHandler) GetAllForRequester(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)

	tags, err := h.TagsService.GetAll(user)

	if err != nil {
		returnErrorResponse(w, http.StatusInternalServerError, errors.New("unable to fetch tags"))
		return
	}
	returnObjectResponse(w, http.StatusCreated, tags)
}
