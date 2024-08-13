package utils

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

var instance *cache.Cache

var cacheCtx = context.Background()

func InitCache() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
	_, err := client.Ping(cacheCtx).Result()
	if err != nil {
		Fatal("Redis connection was refused: %s", err)
	}

	LFUSize, err := strconv.Atoi(os.Getenv("LFU_SIZE"))
	if err != nil {
		LFUSize = 1000
	}

	instance = cache.New(&cache.Options{
		Redis:      client,
		LocalCache: cache.NewTinyLFU(LFUSize, time.Second),
	})

	return client
}

func CreateDuration(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}

func GetCache[T any](key string, payload T) error {
	if os.Getenv("ENV") == "development" {
		return nil
	}

	err := instance.Get(cacheCtx, key, payload)
	if err != nil {
		return err
	}

	return nil
}

func SetCache[T any](key string, payload T, TTL time.Duration) error {
	if os.Getenv("ENV") == "development" {
		return nil
	}

	err := instance.Set(&cache.Item{
		Ctx:   cacheCtx,
		Key:   key,
		Value: payload,
		TTL:   TTL,
	})
	if err != nil {
		return err
	}

	return nil
}
