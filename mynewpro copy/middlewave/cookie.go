package middlewave

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	secretKey = "secret"
)

func SetSession() gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte("secret"))
	return sessions.Sessions("mysession", store)
}

func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get(secretKey)
		if sessionID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "请先登录",
			})
			return
		}
		c.Next()
	}
}

// save session
func SaveSession(c *gin.Context, userID string) {
	session := sessions.Default(c)
	session.Set(secretKey, userID)
	session.Save()

}

// clear session
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// get session
func GetSession(c *gin.Context) string {
	session := sessions.Default(c)
	sessionID := session.Get(secretKey)
	return sessionID.(string)
}

// check session
func ChectSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionID := session.Get(secretKey)
	return sessionID == nil

}
