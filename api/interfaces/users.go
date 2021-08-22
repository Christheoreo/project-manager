package interfaces

import "github.com/Christheoreo/project-manager/dtos"

type IUsersRepository interface {
	Insert(user dtos.NewUserDto) (int, error)
	GetByID(id int) (dtos.UserDto, error)
	GetByEmail(email string) (dtos.UserDto, error)
	GetPassword(authLogin dtos.AuthLoginDto) (string, error)
}
type IUsersService interface {
	ValidateNewUser(newUser dtos.NewUserDto) (errorMessages []string, err error)
	HasEmail(email string) (exists bool)
	Insert(newUser dtos.NewUserDto) (dtos.UserDto, error)
	Get(ID int) (user dtos.UserDto, err error)
	GetByEmail(email string) (user dtos.UserDto, err error)
	ValidateCredentials(authLogin dtos.AuthLoginDto) (jwtToken string, err error)
}
