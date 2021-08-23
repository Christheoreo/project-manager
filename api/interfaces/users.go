package interfaces

import "github.com/Christheoreo/project-manager/models"

type IUsersRepository interface {
	Insert(user models.NewUser) (int, error)
	GetByID(id int) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetPassword(authLogin models.AuthLogin) (string, error)
}
type IUsersService interface {
	HasEmail(email string) (exists bool)
	Insert(newUser models.NewUser) (models.User, error)
	Get(ID int) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	ValidateCredentials(authLogin models.AuthLogin) (jwtToken string, err error)
}
