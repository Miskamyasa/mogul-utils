package middlewares

import (
	"encoding/json"
	"github.com/Miskamyasa/mogul-utils/cache"
	"github.com/Miskamyasa/mogul-utils/response"
	"net/http"
)

func GenerateCacheKey(req *http.Request) string {
	ip := req.RemoteAddr
	path := req.URL.Path
	return "cache:" + ip + ":" + path
}

func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var payload *interface{}
		err := cache.GetCache(GenerateCacheKey(req), &payload)
		if err == nil && payload != nil {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(payload)
			if err != nil {
				return
			}
			return
		}
		next.ServeHTTP(w, req)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				response.SendInternalServerError(w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
