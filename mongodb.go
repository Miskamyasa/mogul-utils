package utils

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoDB  *mongo.Database
	mongoCtx = context.TODO()
)

func InitMongoDB() (context.Context, *mongo.Client) {
	ctx, cancel := context.WithTimeout(mongoCtx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		Fatal("Failed to connect to MongoDB", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		Fatal("Failed to ping MongoDB", err)
	}

	mongoDB = client.Database(os.Getenv("MONGODB_NAME"))
	if mongoDB == nil {
		Fatal("Failed to retrieve a database", err)
	}

	return mongoCtx, client
}

func GetMongoDB() (context.Context, *mongo.Database) {
	return mongoCtx, mongoDB
}
