package repositories

import (
	"context"
	"errors"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/Christheoreo/project-manager/utils"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Implements Interface UsersRepository.
type UsersRepositoryPostgres struct {
	Pool *pgxpool.Pool
}

func (r *UsersRepositoryPostgres) Insert(user dtos.NewUserDto) (id int, err error) {
	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		err = errors.New("could not encrypt password.")
		return
	}
	query := "INSERT INTO users (\"first_name\", \"last_name\", \"email\", \"password\") VALUES ($1,$2,$3,$4) RETURNING id"
	err = r.Pool.QueryRow(context.Background(), query, user.FirstName, user.LastName, user.Email, passwordHash).Scan(&id)
	return
}
func (r UsersRepositoryPostgres) GetByID(id int) (user dtos.UserDto, err error) {
	query := "SELECT \"id\", \"first_name\", \"last_name\", \"email\" from \"users\" where \"id\" = $1"
	err = r.Pool.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	return
}
func (r UsersRepositoryPostgres) GetByEmail(email string) (user dtos.UserDto, err error) {
	query := "SELECT \"id\", \"first_name\", \"last_name\", \"email\" from \"users\" where \"email\" = $1"
	err = r.Pool.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	return
}
func (r UsersRepositoryPostgres) GetPassword(authLogin dtos.AuthLoginDto) (string, error) {
	var passwordHash string
	query := "SELECT \"password\" from \"users\" where \"email\" = $1"
	err := r.Pool.QueryRow(context.Background(), query, authLogin.Email).Scan(&passwordHash)
	return passwordHash, err
}
