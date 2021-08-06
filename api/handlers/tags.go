package handlers

import (
	"encoding/json"
	"errors"
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

func (t *TagsHandler) validateNewTagDto(body io.Reader) (newTag dtos.NewTagDto, errorMessages []string, err error) {

	err = json.NewDecoder(body).Decode(&newTag)
	if err != nil {
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

func (t *TagsHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(dtos.UserDto)
	newTag, errMessages, err := t.validateNewTagDto(r.Body)

	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, errMessages)
		return
	}

	// Check if a user with that email exists

	taken, err := t.TagModel.HasTagBeenTakenByUser(newTag.Name, user.ID)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Could not verify if tag already exists."})
		return
	}

	if taken {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"You already have a tag with that name."})
		return
	}

	insertedId, err := t.TagModel.Insert(newTag, user)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to create user."})
		return
	}

	tag, err := t.TagModel.GetById(insertedId)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to fetch created tag."})
		return
	}

	returnObjectResponse(w, http.StatusCreated, tag)
}
