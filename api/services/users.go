package services

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/interfaces"
	"github.com/Christheoreo/project-manager/utils"
)

// Implements IUserService
type UsersService struct {
	UsersRepository interfaces.IUsersRepository
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s *UsersService) ValidateNewUser(newUser dtos.NewUserDto) (errorMessages []string, err error) {

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

func (s *UsersService) ValidateAuthDto(authLogin dtos.AuthLoginDto) (errorMessages []string, err error) {

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

func (s *UsersService) HasEmail(email string) (exists bool) {
	_, err := s.UsersRepository.GetByEmail(email)
	return err == nil
}
func (s UsersService) Insert(newUser dtos.NewUserDto) (dtos.UserDto, error) {

	var user dtos.UserDto

	exists := s.HasEmail(newUser.Email)
	if exists {
		return user, errors.New("user with email already exists")
	}
	ID, err := s.UsersRepository.Insert(newUser)

	if err != nil {
		return user, err
	}
	user, err = s.UsersRepository.GetByID(ID)
	return user, err
}
func (s UsersService) Get(ID int) (user dtos.UserDto, err error) {
	return s.UsersRepository.GetByID(ID)
}
func (s UsersService) GetByEmail(email string) (user dtos.UserDto, err error) {
	return s.UsersRepository.GetByEmail(email)
}
func (s UsersService) ValidateCredentials(authLogin dtos.AuthLoginDto) (string, int, error) {
	var user dtos.UserDto

	passwordHash, err := s.UsersRepository.GetPassword(authLogin)

	if err != nil {
		return "", http.StatusUnprocessableEntity, errors.New("invalid details")
	}

	matches := utils.CheckPasswordHash(authLogin.Password, passwordHash)

	if !matches {
		return "", http.StatusBadRequest, errors.New("invalid login details")
	}

	user, _ = s.GetByEmail(authLogin.Email)

	jwtToken, err := utils.CreateToken(user.ID)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("could not sign the token")
	}

	return jwtToken, -1, nil
}
