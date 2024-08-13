package utils

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisCtx    = context.TODO()
)

func InitRedis() *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := redisClient.Ping(redisCtx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return redisClient
}

func GetRedis() *redis.Client {
	return redisClient
}
