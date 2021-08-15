package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/interfaces"
)

type (
	AuthHandler struct {
		UsersService interfaces.IUsersService
	}
)

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var authLogin dtos.AuthLoginDto

	err := json.NewDecoder(r.Body).Decode(&authLogin)
	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	errMessages, err := h.UsersService.ValidateAuthDto(authLogin)
	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, errMessages)
		return
	}

	// Check if a user with that email and password exists.
	jwtToken, err := h.UsersService.ValidateCredentials(authLogin)

	if err != nil {
		returnErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	returnObjectResponse(w, http.StatusOK, dtos.JwtResponse{
		AccessToken: jwtToken,
	})

}
