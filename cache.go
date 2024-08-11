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

var ctx = context.TODO()

var cacheTTL = 3600

func InitCache() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		Fatal("Redis connection was refused: %s", err)
	}

	envTTL, err := strconv.Atoi(os.Getenv("CACHE_TTL"))
	if err == nil {
		cacheTTL = envTTL
	}

	instance = cache.New(&cache.Options{
		Redis:      client,
		LocalCache: cache.NewTinyLFU(cacheTTL, time.Second),
	})

	return client
}

func GetCache[T any](key string, payload T) error {
	if os.Getenv("ENV") == "development" {
		return nil
	}

	err := instance.Get(ctx, key, payload)
	if err != nil {
		return err
	}

	return nil
}

func SetCache[T any](key string, payload T) error {
	if os.Getenv("ENV") == "development" {
		return nil
	}

	err := instance.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: payload,
		TTL:   time.Duration(cacheTTL) * time.Second,
	})
	if err != nil {
		return err
	}

	return nil
}
