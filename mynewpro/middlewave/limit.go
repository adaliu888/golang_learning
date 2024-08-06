package middlewave

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	//"github.com/gin-gonic/contrib/ratelimit"
)

// 自定义限流
func RateLimitMiddleware() gin.HandlerFunc {
	limiter := rate.NewLimiter(5, 10) // 每秒5个令牌，桶大小为10
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
