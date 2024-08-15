package middlewave

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthTokenMiddlewave() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从HTTP请求头中获取Token
		tokenString := c.GetHeader("Authorization")
		if len(strings.Split(tokenString, " ")) == 2 {
			token, err := jwt.Parse(strings.Split(tokenString, " ")[1], func(token *jwt.Token) (interface{}, error) {
				// 确保Token是使用HS256签名的
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte("your_secret_key"), nil
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				c.Abort()
				return
			}
			// 将Token的Claims数据添加到Context中，以便后续使用
			c.Set("claims", token.Claims)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format is incorrect"})
			c.Abort()
		}
	}
}
