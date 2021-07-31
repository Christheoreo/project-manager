package handlers

import (
	"encoding/json"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/Christheoreo/project-manager/dtos"
)

func returnStandardResponse(w http.ResponseWriter, status int, messages []string) {
	var standardResponse dtos.StandardResponseDto = dtos.StandardResponseDto{
		Status:   status,
		Messages: messages,
	}
	json.NewEncoder(w).Encode(standardResponse)
	w.WriteHeader(status)
}

func returnObjectResponse(w http.ResponseWriter, status int, object interface{}) {
	json.NewEncoder(w).Encode(object)
	w.WriteHeader(status)
}

func isEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}
