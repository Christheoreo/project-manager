package services

import (
	"errors"
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
func (s UsersService) ValidateCredentials(authLogin dtos.AuthLoginDto) (string, error) {
	var user dtos.UserDto

	passwordHash, err := s.UsersRepository.GetPassword(authLogin)

	if err != nil {
		return "", errors.New("invalid details")
	}

	matches := utils.CheckPasswordHash(authLogin.Password, passwordHash)

	if !matches {
		return "", errors.New("invalid login details")
	}

	user, _ = s.GetByEmail(authLogin.Email)

	jwtToken, err := utils.CreateToken(user.ID)

	if err != nil {
		return "", errors.New("could not sign the token")
	}

	return jwtToken, nil
}
