package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key") // 用于签名的密钥

// Claims 定义 JWT 的载荷
type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

// GenerateJWT 生成 JWT
func GenerateJWT(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // 设置过期时间
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// ValidateJWT 验证 JWT
func ValidateJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        return nil, err
    }
    return claims, nil
}