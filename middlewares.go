package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func GenerateCacheKey(req *fiber.Ctx) string {
	return "cache:" + req.IP() + ":" + req.Path()
}

func CacheMiddleware(req *fiber.Ctx) error {
	var payload *interface{}
	err := GetCache(GenerateCacheKey(req), &payload)
	if err == nil && payload != nil {
		return req.JSON(payload)
	}
	return req.Next()
}

func RateLimiterMiddleware() fiber.Handler {
	var Next func(req *fiber.Ctx) bool

	if os.Getenv("ENV") == "development" {
		Next = func(req *fiber.Ctx) bool {
			return req.IP() == "127.0.0.1"
		}
	} else {
		Next = limiter.ConfigDefault.Next
	}

	limiterTTL, err := time.ParseDuration(os.Getenv("RATE_LIMITER_TTL"))
	if err != nil {
		limiterTTL = 5 * time.Minute
	}

	return limiter.New(limiter.Config{
		Next:       Next,
		Max:        1,
		Expiration: limiterTTL,
	})
}
