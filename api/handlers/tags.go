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

func (h *TagsHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)
	var newTag dtos.NewTagDto

	err := json.NewDecoder(r.Body).Decode(&newTag)
	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	errMessages, err := h.TagsService.ValidateNewTagDto(newTag)

	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, errMessages)
		return
	}

	tag, err := h.TagsService.Create(newTag, user)
	if err != nil {
		returnErrorResponse(w, http.StatusInternalServerError, err)
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
