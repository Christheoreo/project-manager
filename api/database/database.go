package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EstablishClient() (*mongo.Client, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connectionURL := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUsername, dbPassword, dbHost, dbPort)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURL))
	if err != nil {
		return client, err
	}
	// this ctx is only for the connection - different contexts can be made for requests.
	/**
	 * The ten seconds I used might be a little too generous for your needs,
	 * but feel free to play around with the value that makes the most sense to you.
	 */
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	return client, err
}

func EndClient(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Disconnect(ctx)
}
