package repositories

import (
	"context"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Implements Interface UsersRepository.
type UsersRepositoryPostgres struct {
	Pool *pgxpool.Pool
}

func (r *UsersRepositoryPostgres) Insert(user dtos.NewUserDto) (id int, err error) {
	query := `INSERT INTO users ("first_name", "last_name", "email", "password") VALUES ($1,$2,$3,$4) RETURNING id`
	err = r.Pool.QueryRow(context.Background(), query, user.FirstName, user.LastName, user.Email, user.Password).Scan(&id)
	return
}
func (r UsersRepositoryPostgres) GetByID(id int) (user dtos.UserDto, err error) {
	query := `SELECT "id", "first_name", "last_name", "email" FROM "users" WHERE "id" = $1`
	err = r.Pool.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	return
}
func (r UsersRepositoryPostgres) GetByEmail(email string) (user dtos.UserDto, err error) {
	query := `SELECT "id", "first_name", "last_name", "email" FROM "users" WHERE "email" = $1`
	err = r.Pool.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	return
}
func (r UsersRepositoryPostgres) GetPassword(authLogin dtos.AuthLoginDto) (string, error) {
	var passwordHash string
	query := `SELECT "password" FROM "users" WHERE "email" = $1`
	err := r.Pool.QueryRow(context.Background(), query, authLogin.Email).Scan(&passwordHash)
	return passwordHash, err
}
