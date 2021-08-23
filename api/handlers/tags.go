package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Christheoreo/project-manager/interfaces"
	"github.com/Christheoreo/project-manager/models"
)

type (
	TagsHandler struct {
		TagsService interfaces.ITagsService
	}
)

func (h *TagsHandler) validateNewTagDto(newTag models.NewTag) ([]string, error) {

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
	user := getUser(r)
	var newTag models.NewTag

	err := json.NewDecoder(r.Body).Decode(&newTag)
	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	errMessages, err := h.validateNewTagDto(newTag)

	if err != nil {
		returnStandardResponse(w, http.StatusUnprocessableEntity, errMessages)
		return
	}

	tag, err := h.TagsService.Create(newTag, user)
	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	returnObjectResponse(w, http.StatusCreated, tag)
}

func (h *TagsHandler) GetAllForRequester(w http.ResponseWriter, r *http.Request) {
	user := getUser(r)

	tags, err := h.TagsService.GetAll(user)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"unable to fetch tags"})
		return
	}
	returnObjectResponse(w, http.StatusCreated, tags)
}
