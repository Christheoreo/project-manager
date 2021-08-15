package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Implements Interface PrioritiesRepository.
type PrioritiesRepositoryPostgres struct {
	Pool *pgxpool.Pool
}

func (r *PrioritiesRepositoryPostgres) Exists(ID int) (exists bool, err error) {
	var count int
	query := "SELECT COUNT(*) FROM \"priorities\" where \"id\" = $1"
	err = r.Pool.QueryRow(context.Background(), query, ID).Scan(&count)
	if err != nil {
		return
	}
	exists = count > 0
	return
}
