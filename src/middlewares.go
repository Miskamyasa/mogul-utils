package utils

import (
	"encoding/json"
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
		err := GetCache(GenerateCacheKey(req), &payload)
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
