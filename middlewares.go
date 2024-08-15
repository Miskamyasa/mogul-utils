package utils

import (
	"github.com/gofiber/fiber/v2"
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
