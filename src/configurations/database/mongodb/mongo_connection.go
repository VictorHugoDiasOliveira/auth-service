package mongodb

import (
	"context"
	"os"
	"sosservice/src/configurations/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	MONGODB_URL      = "MONGODB_URL"
	MONGODB_DATABASE = "MONGODB_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	logger.Info("Starting Database Connection", zap.String("journey", "NewMongoDBConnection"))
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_DATABASE)

	// TODO Def password
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		logger.Error("Error trying to start Database", err, zap.String("journey", "NewMongoDBConnection"))
		return nil, err
	}
	logger.Info("Database Connection Success", zap.String("journey", "NewMongoDBConnection"))
	return client.Database(mongodb_database), nil
}
