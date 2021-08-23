package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Christheoreo/project-manager/interfaces"
	"github.com/Christheoreo/project-manager/models"
)

type (
	AuthHandler struct {
		UsersService interfaces.IUsersService
	}
)

func (h *AuthHandler) validateAuthDto(authLogin models.AuthLogin) (errorMessages []string, err error) {

	errorMessages = make([]string, 0)

	if !isEmailValid(authLogin.Email) {
		e := "'email' needs to be valid"
		errorMessages = append(errorMessages, e)
	}

	if len(authLogin.Password) < 8 {
		e := "'password' needs to be at least 8 characters"
		errorMessages = append(errorMessages, e)
	}

	if len(errorMessages) > 0 {
		err = errors.New("invalid DTO")
	}
	return
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var authLogin models.AuthLogin

	err := json.NewDecoder(r.Body).Decode(&authLogin)
	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	errMessages, err := h.validateAuthDto(authLogin)
	if err != nil {
		returnStandardResponse(w, http.StatusUnprocessableEntity, errMessages)
		return
	}

	// Check if a user with that email and password exists.
	jwtToken, err := h.UsersService.ValidateCredentials(authLogin)

	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	returnObjectResponse(w, http.StatusOK, models.JwtResponse{
		AccessToken: jwtToken,
	})

}
