package models

import (
	"context"
	"fmt"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/utils"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	Pool *pgxpool.Pool
}

func (u *User) HasEmailBeenTaken(email string) (bool, error) {
	var count int
	err := u.Pool.QueryRow(context.Background(), "SELECT COUNT(*) as \"count\" from \"users\" where \"email\" = $1", email).Scan(&count)
	return count > 0, err

}

func (u *User) Insert(user dtos.NewUserDto) (id int, err error) {
	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return
	}
	query := "INSERT INTO users (\"first_name\", \"last_name\", \"email\", \"password\") VALUES ($1,$2,$3,$4) RETURNING id"
	err = u.Pool.QueryRow(context.Background(), query, user.FirstName, user.LastName, user.Email, passwordHash).Scan(&id)
	return
}

func (u *User) GetById(id int) (user dtos.UserDto, err error) {
	fmt.Printf("HELLO CHRIS ID = %d\n\n\n", id)
	query := "SELECT \"id\", \"first_name\", \"last_name\", \"email\" from \"users\" where \"id\" = $1"
	err = u.Pool.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	return
}

func (u *User) GetByEmail(email string) (user dtos.UserDto, err error) {
	query := "SELECT \"id\", \"first_name\", \"last_name\", \"email\" from \"users\" where \"email\" = $1"
	err = u.Pool.QueryRow(context.Background(), query, email).Scan(&user)
	return
}

func (u *User) ValidateUserCredentials(authLogin dtos.AuthLoginDto) (valid bool, err error) {

	var passwordHash string
	query := "SELECT \"password\" from \"users\" where \"email\" = $1"
	err = u.Pool.QueryRow(context.Background(), query, authLogin.Password).Scan(&passwordHash)
	if err != nil {
		return
	}
	valid = utils.CheckPasswordHash(authLogin.Password, passwordHash)
	return
}
