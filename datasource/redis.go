package datasource

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisCtx    = context.Background()
)

func InitRedis() *redis.Client {
	config := &redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	}

	if os.Getenv("REDIS_PASSWORD") != "" {
		config.Password = os.Getenv("REDIS_PASSWORD")
	}

	redisClient = redis.NewClient(config)

	_, err := redisClient.Ping(redisCtx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return redisClient
}

func GetRedis() (context.Context, *redis.Client) {
	return redisCtx, redisClient
}
