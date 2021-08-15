package handlers

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/Christheoreo/project-manager/dtos"
)

func returnStandardResponse(w http.ResponseWriter, status int, messages []string) {
	var standardResponse dtos.StandardResponseDto = dtos.StandardResponseDto{
		Status:   status,
		Messages: messages,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(standardResponse)
}

func returnErrorResponse(w http.ResponseWriter, status int, err error) {
	var standardResponse dtos.StandardResponseDto = dtos.StandardResponseDto{
		Status:   status,
		Messages: []string{err.Error()},
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(standardResponse)
}

func returnObjectResponse(w http.ResponseWriter, status int, object interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(object)
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
