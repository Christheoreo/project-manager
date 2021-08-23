package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Implements Interface PrioritiesRepository.
type PrioritiesRepositoryPostgres struct {
	Pool *pgxpool.Pool
}

func (r *PrioritiesRepositoryPostgres) Exists(ID int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM "priorities" where "id" = $1`
	if err := r.Pool.QueryRow(context.Background(), query, ID).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
