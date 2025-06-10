package middlewave

import (
	"net/http"
	"strings"

	"golang_learning/mynewpro/jwt"

	"github.com/gin-gonic/gin"
)

// JWT认证中间件 - 验证请求中的令牌
func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从HTTP请求头中获取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}

		// 检查bearer格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be 'Bearer {token}'"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 使用JWT包验证令牌
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// 将用户ID添加到上下文
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

// 刷新令牌处理函数 - 可用于创建刷新令牌API端点
func RefreshTokenHandler(c *gin.Context) {
	// 从请求中获取刷新令牌
	refreshToken := c.GetHeader("Refresh-Token")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token is required"})
		return
	}

	// 使用刷新令牌获取新的访问令牌
	newToken, err := jwt.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newToken,
	})
}

// 注销处理函数 - 撤销令牌并清除刷新令牌
func LogoutHandler(c *gin.Context) {
	// 获取访问令牌
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" && len(strings.Split(authHeader, " ")) == 2 {
		tokenString := strings.Split(authHeader, " ")[1]

		// 验证令牌获取用户ID
		claims, err := jwt.ValidateToken(tokenString)
		if err == nil {
			// 撤销访问令牌
			_ = jwt.RevokeToken(tokenString)

			// 清除用户的所有刷新令牌
			_ = jwt.ClearUserRefreshTokens(claims.UserID)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
