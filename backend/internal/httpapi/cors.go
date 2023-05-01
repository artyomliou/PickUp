package httpapi

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOriginFunc = func(origin string) bool {
		return origin == "null" || strings.Contains(origin, "://localhost") || strings.Contains(origin, "://127.0.0.1") // TODO for local development
	}
	config.AllowCredentials = true
	// config.MaxAge = 12 * time.Hour
	return cors.New(config)
}
