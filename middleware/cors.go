package middleware

import (
	"chat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var AllowedOrigins = []string{
	"https://fystart.cn",
	"https://www.fystart.cn",
	"https://nio.fystart.cn",
	"https://chatnio.net",
	"https://www.chatnio.net",
	"http://localhost:5173",
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if utils.Contains(origin, AllowedOrigins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request.Method == "OPTIONS" {
				c.Writer.Header().Set("Access-Control-Max-Age", "7200")
				c.AbortWithStatus(http.StatusOK)
				return
			}
		}

		c.Next()
	}
}
