package handlers

import (
	"encoding/json"
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

func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var newUser dtos.NewUserDto
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	errMessages, err := h.UsersService.ValidateNewUser(newUser)
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
