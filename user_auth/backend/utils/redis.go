package utils

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

// InitRedis 初始化 Redis 客户端
func InitRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis 地址
    })

    // 测试连接
    _, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        Logger.Fatal("Failed to connect to Redis:", err)
    }
}