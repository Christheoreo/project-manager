package models

import (
	"context"

	"github.com/Christheoreo/project-manager/dtos"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Tag struct {
	Pool *pgxpool.Pool
}

func (t *Tag) Insert(tag dtos.NewTagDto, userId int) (id int, err error) {

	query := "INSERT INTO tags (\"name\", \"user_id\") VALUES ($1,$2) RETURNING id"
	err = t.Pool.QueryRow(context.Background(), query, tag.Name, userId).Scan(&id)
	return
}

/**
@todo this
**/
// func (u *User) GetById(id int) (user dtos.UserDto, err error) {
// 	fmt.Printf("HELLO CHRIS ID = %d\n\n\n", id)
// 	query := "SELECT \"id\", \"first_name\", \"last_name\", \"email\" from \"users\" where \"id\" = $1"
// 	err = u.Pool.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
// 	return
// }
