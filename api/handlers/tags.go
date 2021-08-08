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
	TagsHandler struct {
		TagModel models.Tag
	}
)

func (h *TagsHandler) validateNewUserDto(body io.Reader) (newTag dtos.NewTagDto, errorMessages []string, err error) {

	err = json.NewDecoder(body).Decode(&newTag)
	if err != nil {
		fmt.Println("Got here")
		return
	}

	errorMessages = make([]string, 0)

	if len(newTag.Name) < 3 {
		e := "'name' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}

	if len(errorMessages) > 0 {
		err = errors.New("invalid DTO")
	}
	return
}

func (h *TagsHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)
	newTag, errMessages, err := h.validateNewUserDto(r.Body)
	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, errMessages)
		return
	}

	// Check if that tag exists for this user.

	taken, err := h.TagModel.DoesTagExistForUser(newTag.Name, user.ID)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Could not verify if tag already exists."})
		return
	}

	if taken {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"You already have a tag with this name."})
		return
	}

	insertedId, err := h.TagModel.Insert(newTag, user.ID)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to create tag."})
		return
	}

	tag, err := h.TagModel.GetById(insertedId)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to fetch created tag."})
		return
	}

	returnObjectResponse(w, http.StatusCreated, tag)
}

func (h *TagsHandler) GetAllForRequester(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)

	tags, err := h.TagModel.GetAllForUser(user.ID)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to fetch tags."})
		return
	}

	returnObjectResponse(w, http.StatusCreated, tags)
}
