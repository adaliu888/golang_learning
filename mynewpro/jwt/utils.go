package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secret = []byte("your_secret_key")
)

type Claims struct {
	UserID string
	jwt.StandardClaims
}

// 新建一个token
func GenerateToken(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		// 设置Token的过期时间
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			Issuer:    "gin-example",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// token 保存
/*func SaveTokenToRedis(token string, userId uint) error {
	cacheKey := fmt.Sprintf("user:%d:token", userId)
	return global.App.Redis.SetEX(context.Background(), cacheKey, token, time.Duration(global.App.Config.Jwt.JwtTtl)*time.Second).Err()
}
*/
