package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func EstablishConnectionPool() (*pgxpool.Pool, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// return pgx.Connect(context.Background(), connectionURL)
	return pgxpool.Connect(context.Background(), connectionURL)
}
func CloseConnectionPool(conn *pgxpool.Pool) {
	conn.Close()
}
