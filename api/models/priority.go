package models

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Priority struct {
		Pool *pgxpool.Pool
	}
)

func (p *Priority) Exists(priorityID int) (exists bool, err error) {
	var count int
	query := "SELECT COUNT(*) FROM \"priorities\" where \"id\" = $1"
	err = p.Pool.QueryRow(context.Background(), query, priorityID).Scan(&count)
	if err != nil {
		return
	}
	exists = count > 0
	return
}
