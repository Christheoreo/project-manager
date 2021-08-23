package repositories

import (
	"context"

	"github.com/Christheoreo/project-manager/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Implements Interface TagsRepository.
type TagsRepositoryPostgres struct {
	Pool *pgxpool.Pool
}

func (r *TagsRepositoryPostgres) Insert(tag models.NewTagDto, userId int) (id int, err error) {
	query := `INSERT INTO tags ("name", user_id) VALUES ($1,$2) RETURNING id`
	err = r.Pool.QueryRow(context.Background(), query, tag.Name, userId).Scan(&id)
	return
}

func (r *TagsRepositoryPostgres) GetById(id int) (tag models.TagDto, err error) {
	query := `SELECT id, "name" FROM tags WHERE id = $1`
	err = r.Pool.QueryRow(context.Background(), query, id).Scan(&tag.ID, &tag.Name)
	return
}
func (r *TagsRepositoryPostgres) GetAllForUser(userID int) (tags []models.TagDto, err error) {
	query := `SELECT id, "name" from tags where user_id = $1`
	rows, err := r.Pool.Query(context.Background(), query, userID)
	if err != nil {
		return
	}
	tags = make([]models.TagDto, 0)
	for rows.Next() {
		var tag models.TagDto

		err = rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return
		}
		tags = append(tags, tag)
	}
	return
}
func (r *TagsRepositoryPostgres) Exists(name string, userID int) (exists bool, err error) {
	var count int
	exists = true
	query := `SELECT COUNT(*) as "count" from tags WHERE "name" = $1 AND user_id = $2`
	err = r.Pool.QueryRow(context.Background(), query, name, userID).Scan(&count)
	if err != nil {
		return
	}
	exists = count > 0
	return
}
