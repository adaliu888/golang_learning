package jwt

import (
	"errors"
	"fmt"
	"time"

	"golang_learning/mynewpro/db"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

// JWTConfig 存储JWT配置
type JWTConfig struct {
	Secret        string        `mapstructure:"secret"`
	TokenExpiry   time.Duration `mapstructure:"token_expiry"`
	RefreshExpiry time.Duration `mapstructure:"refresh_expiry"`
}

// 全局配置
var (
	jwtConfig JWTConfig
)

// Claims 是JWT的声明部分
type Claims struct {
	UserID string
	jwt.StandardClaims
}

// 初始化JWT配置
func InitJWTConfig() error {
	// 使用viper从配置文件加载
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 如果配置文件中没有JWT相关配置，使用默认值
	if !v.IsSet("jwt") {
		jwtConfig = JWTConfig{
			Secret:        "your_secret_key_change_in_production",
			TokenExpiry:   15 * time.Minute,
			RefreshExpiry: 24 * time.Hour,
		}
		return nil
	}

	if err := v.UnmarshalKey("jwt", &jwtConfig); err != nil {
		return fmt.Errorf("解析JWT配置失败: %v", err)
	}

	return nil
}

// 生成JWT令牌
func GenerateToken(userID string) (string, error) {
	// 确保配置已初始化
	if jwtConfig.Secret == "" {
		if err := InitJWTConfig(); err != nil {
			return "", err
		}
	}

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtConfig.TokenExpiry).Unix(),
			Issuer:    "gin-example",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}

// 生成刷新令牌
func GenerateRefreshToken(userID string) (string, error) {
	// 确保配置已初始化
	if jwtConfig.Secret == "" {
		if err := InitJWTConfig(); err != nil {
			return "", err
		}
	}

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtConfig.RefreshExpiry).Unix(),
			Issuer:    "gin-example-refresh",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return "", err
	}

	// 存储刷新令牌到Redis
	if err := StoreRefreshToken(userID, refreshToken, jwtConfig.RefreshExpiry); err != nil {
		return "", err
	}

	return refreshToken, nil
}

// 验证令牌
func ValidateToken(tokenString string) (*Claims, error) {
	// 确保配置已初始化
	if jwtConfig.Secret == "" {
		if err := InitJWTConfig(); err != nil {
			return nil, err
		}
	}

	// 检查令牌是否被撤销
	if IsTokenRevoked(tokenString) {
		return nil, errors.New("令牌已被撤销")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("非预期的签名方法: %v", token.Header["alg"])
		}
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("无效的令牌")
	}

	return claims, nil
}

// 使用刷新令牌获取新令牌
func RefreshToken(refreshToken string) (string, error) {
	claims, err := ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	// 验证刷新令牌是否存在于Redis中
	exists, err := ValidateRefreshTokenInRedis(claims.UserID, refreshToken)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", errors.New("刷新令牌无效或已过期")
	}

	// 生成新的访问令牌
	return GenerateToken(claims.UserID)
}

// 撤销令牌 - 将令牌加入黑名单
func RevokeToken(tokenString string) error {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return err
	}

	// 计算令牌的剩余有效期
	expirationTime := time.Unix(claims.ExpiresAt, 0)
	ttl := time.Until(expirationTime)

	if ttl <= 0 {
		// 令牌已过期，无需加入黑名单
		return nil
	}

	// 将令牌添加到Redis黑名单
	conn := db.Redisdefaultpool.Get()
	defer conn.Close()

	// 使用令牌本身作为键，并设置与令牌相同的过期时间
	_, err = conn.Do("SETEX", "revoked_token:"+tokenString, int(ttl.Seconds()), 1)
	return err
}

// 检查令牌是否被撤销
func IsTokenRevoked(tokenString string) bool {
	conn := db.Redisdefaultpool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", "revoked_token:"+tokenString))
	if err != nil {
		// 出错时，为安全起见，视为已撤销
		return true
	}

	return exists
}

// 存储刷新令牌到Redis
func StoreRefreshToken(userID string, refreshToken string, expiry time.Duration) error {
	conn := db.Redisdefaultpool.Get()
	defer conn.Close()

	// 存储刷新令牌，使用用户ID作为键的一部分
	key := fmt.Sprintf("refresh_token:%s", userID)
	_, err := conn.Do("SETEX", key, int(expiry.Seconds()), refreshToken)
	return err
}

// 验证刷新令牌是否存在于Redis中
func ValidateRefreshTokenInRedis(userID string, refreshToken string) (bool, error) {
	conn := db.Redisdefaultpool.Get()
	defer conn.Close()

	key := fmt.Sprintf("refresh_token:%s", userID)
	storedToken, err := redis.String(conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			// 键不存在
			return false, nil
		}
		return false, err
	}

	// 验证存储的令牌是否与提供的令牌匹配
	return storedToken == refreshToken, nil
}

// 清除用户的所有刷新令牌（用于注销）
func ClearUserRefreshTokens(userID string) error {
	conn := db.Redisdefaultpool.Get()
	defer conn.Close()

	key := fmt.Sprintf("refresh_token:%s", userID)
	_, err := conn.Do("DEL", key)
	return err
}
