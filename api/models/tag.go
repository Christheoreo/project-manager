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

func (t *Tag) GetById(id int) (tag dtos.TagDto, err error) {
	query := "SELECT \"id\", \"name\" from \"tags\" where \"id\" = $1"
	err = t.Pool.QueryRow(context.Background(), query, id).Scan(&tag.ID, &tag.Name)
	return
}

func (t *Tag) GetAllForUser(userID int) (tags []dtos.TagDto, err error) {
	query := "SELECT \"id\", \"name\" from \"tags\" where \"user_id\" = $1"
	rows, err := t.Pool.Query(context.Background(), query, userID)

	if err != nil {
		return
	}

	tags = make([]dtos.TagDto, 0)

	for rows.Next() {
		tag := dtos.TagDto{
			ID:   0,
			Name: "NA",
		}
		err = rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return
		}
		tags = append(tags, tag)
	}
	return
}

func (t *Tag) DoesTagExistForUser(name string, userID int) (exists bool, err error) {
	var count int
	query := "SELECT COUNT(*) as \"count\" from \"tags\" where \"name\" = $1 AND \"user_id\" = $2"
	err = t.Pool.QueryRow(context.Background(), query, name, userID).Scan(&count)
	if err == nil {
		exists = count > 0
	}
	return
}
