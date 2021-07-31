package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/models"
)

type (
	UserHandler struct {
		UserModel models.User
	}
)

func (h *UserHandler) validateNewUserDto(body io.Reader) (newUser dtos.NewUserDto, errorMessages []string, err error) {

	err = json.NewDecoder(body).Decode(&newUser)
	if err != nil {
		fmt.Println("Got here")
		return
	}

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
	w.Header().Set("Content-Type", "application/json")

	newUser, errMessages, err := h.validateNewUserDto(r.Body)

	if err != nil {
		var standardResponse dtos.StandardResponseDto = dtos.StandardResponseDto{
			Status:   http.StatusBadRequest,
			Messages: errMessages,
		}
		json.NewEncoder(w).Encode(standardResponse)
		return
	}

	// Check if a user with that email exists

	taken, err := h.UserModel.HasEmailBeenTaken(newUser.Email)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Could not verify if email was taken."})
		return
	}

	if taken {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"The email is already in use."})
		return
	}

	insertedId, err := h.UserModel.Insert(newUser)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to create user."})
		return
	}

	user, err := h.UserModel.GetById(insertedId)

	if err != nil {
		returnStandardResponse(w, http.StatusInternalServerError, []string{"Unable to fetch created user."})
		return
	}

	returnObjectResponse(w, http.StatusCreated, user)

}
