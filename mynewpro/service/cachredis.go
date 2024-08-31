package service

import (
	"fmt"

	"net/http"

	red "golang_learning/mynewpro/db"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
)

func CacheOneUserDecorator(h gin.HandlerFunc, pora string, readKeyPattern string, empty interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设 URL 参数中有一个名为 "id" 的参数
		keyId := c.Param(pora)

		// 使用提供的模式和 keyId 生成 Redis 键
		redisKey := fmt.Sprintf(readKeyPattern, keyId)

		// 输出 keyId 和 redisKey 到控制台，实际使用中可能不需要这行
		//fmt.Println(keyId, redisKey)

		// 连接到 Redis
		conn := red.Redisdefaultpool.Get()
		defer conn.Close()
		data, err := redis.Bytes(conn.Do("GET", redisKey))

		// 检查 Redis 是否返回了数据
		if err != nil {
			h(c)
			dbResult, exists := c.Get("dbResult")
			if !exists {
				dbResult = empty
			}
			redisData, _ := ffjson.Marshal(dbResult)

			// 存入 Redis，并返回数据
			conn.Do("SETEX", redisKey, 30, redisData)
			c.JSON(http.StatusOK, gin.H{
				"message": "from db",
				"data":    dbResult,
			})
			return
		}
		// 反序列化 JSON 并返回数据

		ffjson.Unmarshal(data, &empty)
		c.JSON(http.StatusOK, gin.H{
			"message": "from redis",
			"data":    empty,
		})

	}
}
