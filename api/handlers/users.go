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
	UserHandler struct {
		UsersService interfaces.IUsersService
	}
)

func (h *UserHandler) validateNewUser(newUser dtos.NewUserDto) (errorMessages []string, err error) {

	errorMessages = make([]string, 0)

	if len(newUser.FirstName) < 3 {
		e := "'firstName' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}
	if len(newUser.LastName) < 3 {
		e := "'lastName' needs to be at least 3 characters"
		errorMessages = append(errorMessages, e)
	}

	if !isEmailValid(newUser.Email) {
		e := "'email' needs to be valid"
		errorMessages = append(errorMessages, e)
	}

	if len(newUser.Password) < 8 {
		e := "'password' needs to be at least 8 characters"
		errorMessages = append(errorMessages, e)
	}

	if len(newUser.Password) > 25 {
		e := "'password' needs to be less than or equal to 25 characters"
		errorMessages = append(errorMessages, e)
	}

	if len(newUser.Password) >= 8 && newUser.Password != newUser.PasswordConfirm {
		e := "the passwords do not match"
		errorMessages = append(errorMessages, e)
	}

	if len(errorMessages) > 0 {
		err = errors.New("invalid DTO")
	}
	return
}

func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var newUser dtos.NewUserDto
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	errMessages, err := h.validateNewUser(newUser)
	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, errMessages)
		return
	}

	insertedUser, err := h.UsersService.Insert(newUser)
	if err != nil {
		returnErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	returnObjectResponse(w, http.StatusCreated, insertedUser)
}

func (h *UserHandler) GetRequester(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey)
	returnObjectResponse(w, http.StatusOK, user)
}
