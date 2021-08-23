package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Christheoreo/project-manager/models"
)

func returnStandardResponse(w http.ResponseWriter, status int, messages []string) {
	var standardResponse models.StandardResponse = models.StandardResponse{
		Status:   status,
		Messages: messages,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(standardResponse)
}

func returnErrorResponse(w http.ResponseWriter, status int, err error) {
	var standardResponse models.StandardResponse = models.StandardResponse{
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

func getUser(r *http.Request) models.User {
	return r.Context().Value(ContextUserKey).(models.User)
}
