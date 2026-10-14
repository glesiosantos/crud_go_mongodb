package database

import (
	"context"
	"os"

	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGO_DB_URL  = "MONGO_DB_URL"
	MONGO_DB_NAME = "MONGO_DB_NAME"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongoDbUrl := os.Getenv(MONGO_DB_URL)
	mongoDbName := os.Getenv(MONGO_DB_NAME)
	client, err := mongo.Connect(options.Client().ApplyURI(mongoDbUrl))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Database is running")
	return client.Database(mongoDbName), nil
}
