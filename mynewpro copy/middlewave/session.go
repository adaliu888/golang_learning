package middlewave

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/sessions"
)

// 用户登录权限，可以访问其他路由
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//before request
		session := sessions.Default(c)
		sessionID := session.Get("userID")
		if sessionID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		userId := sessionID.(string)
		if userId == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("userID", userId)

		//after request
		c.Next()

	}

}
