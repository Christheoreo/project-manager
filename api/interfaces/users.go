package interfaces

import "github.com/Christheoreo/project-manager/models"

type IUsersRepository interface {
	Insert(user models.NewUserDto) (int, error)
	GetByID(id int) (models.UserDto, error)
	GetByEmail(email string) (models.UserDto, error)
	GetPassword(authLogin models.AuthLoginDto) (string, error)
}
type IUsersService interface {
	HasEmail(email string) (exists bool)
	Insert(newUser models.NewUserDto) (models.UserDto, error)
	Get(ID int) (user models.UserDto, err error)
	GetByEmail(email string) (user models.UserDto, err error)
	ValidateCredentials(authLogin models.AuthLoginDto) (jwtToken string, err error)
}
