package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/Christheoreo/project-manager/interfaces"
	"github.com/Christheoreo/project-manager/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Implements IUserService
type UsersService struct {
	UsersRepository interfaces.IUsersRepository
}

var signingKey []byte

func init() {
	jwtKey := os.Getenv("JWT_KEY")
	signingKey = []byte(jwtKey)
}

func createJwtToken(userId int) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims =
		&jwt.StandardClaims{
			Subject:   strconv.Itoa(userId),
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		}

	return token.SignedString(signingKey)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UsersService) HasEmail(email string) (exists bool) {
	_, err := s.UsersRepository.GetByEmail(email)
	return err == nil
}
func (s UsersService) Insert(newUser models.NewUser) (models.User, error) {

	var user models.User

	exists := s.HasEmail(newUser.Email)
	if exists {
		return user, errors.New("user with email already exists")
	}

	passwordHash, err := hashPassword(newUser.Password)
	if err != nil {
		err = errors.New("could not encrypt password")
		return user, err
	}

	newUser.Password = passwordHash
	newUser.PasswordConfirm = passwordHash

	ID, err := s.UsersRepository.Insert(newUser)

	if err != nil {
		return user, err
	}
	user, err = s.UsersRepository.GetByID(ID)
	return user, err
}
func (s UsersService) Get(ID int) (user models.User, err error) {
	return s.UsersRepository.GetByID(ID)
}
func (s UsersService) GetByEmail(email string) (user models.User, err error) {
	return s.UsersRepository.GetByEmail(email)
}
func (s UsersService) ValidateCredentials(authLogin models.AuthLogin) (string, error) {
	var user models.User

	passwordHash, err := s.UsersRepository.GetPassword(authLogin)

	if err != nil {
		return "", errors.New("invalid details")
	}

	matches := checkPasswordHash(authLogin.Password, passwordHash)

	if !matches {
		return "", errors.New("invalid login details")
	}

	user, _ = s.GetByEmail(authLogin.Email)

	jwtToken, err := createJwtToken(user.ID)

	if err != nil {
		return "", errors.New("could not sign the token")
	}

	return jwtToken, nil
}
