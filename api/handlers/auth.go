package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/models"
	"github.com/Christheoreo/project-manager/utils"
)

type (
	AuthHandler struct {
		UserModel models.User
	}
)

func (h *AuthHandler) validateAuthLoginDto(body io.Reader) (authLogin dtos.AuthLoginDto, errorMessages []string, err error) {

	err = json.NewDecoder(body).Decode(&authLogin)
	if err != nil {
		fmt.Println("Got here")
		return
	}

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
	authLogin, errMessages, err := h.validateAuthLoginDto(r.Body)

	if err != nil {
		returnStandardResponse(w, http.StatusBadRequest, errMessages)
		return
	}

	// Check if a user with that email and password exists.

	isValid, err := h.UserModel.ValidateUserCredentials(authLogin)

	// We want to give a generic error message, so a hacker can't tell if that user exists.
	if err != nil || !isValid {
		returnStandardResponse(w, http.StatusBadRequest, []string{"We could not validate your details."})
		return
	}

	user, _ := h.UserModel.GetByEmail(authLogin.Email)

	jwtToken, err := utils.CreateToken(user.ID)

	if err != nil {
		fmt.Println(err.Error())
		returnStandardResponse(w, http.StatusInternalServerError, []string{"We could not sign the token."})
		return
	}

	returnObjectResponse(w, http.StatusOK, dtos.JwtResponse{
		AccessToken: jwtToken,
	})

}
