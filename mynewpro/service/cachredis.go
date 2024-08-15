package service

import (
	"encoding/json"
	"fmt"
	red "golang_learning/mynewpro/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
)

//cach redis
/*
func CachOneUserDecoractor(h gin.HandlerFunc, porm string, readKeyPattern string, empty interface{}) {
	return func(c *gin.Context) {
		keyId := c.Param(porm)
		redisKey := fmt.Sprintf(readKeyPattern, keyId)
		conn := red.RedisDefaultpool.get()
		defer conn.close()

	}

}*/

/*
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

		err = json.Unmarshal(data, &empty)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "from redis",
			"data":    empty,
		})

	}
}*/

func CacheOneUserDecorator(h gin.HandlerFunc, id string, readKeyPattern string, emptyPtr interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		keyId := c.Param(id)
		redisKey := fmt.Sprintf(readKeyPattern, keyId)
		conn := red.Redisdefaultpool.Get()
		defer conn.Close()

		data, err := redis.Bytes(conn.Do("GET", redisKey))
		if err != nil && err != redis.ErrNil {
			// Redis 错误处理
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get data from Redis"})
			return
		} else if err == nil {
			// 缓存中没有数据，从数据库获取
			h(c)
			dbResult, exists := c.Get("dbResult")
			if !exists {
				// 如果数据库中也没有数据，则返回空
				c.JSON(http.StatusOK, gin.H{
					"message": "not found",
					"data":    emptyPtr,
				})
				return
			}

			// 序列化数据库结果
			redisData, err := ffjson.Marshal(dbResult)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize data"})
				return
			}

			// 存入 Redis，并设置过期时间为 30 秒
			_, err = conn.Do("SETEX", redisKey, 30, redisData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set data in Redis"})
				return
			}

			// 返回数据库查询结果
			c.JSON(http.StatusOK, gin.H{
				"message": "from db",
				"data":    dbResult,
			})
		} else {
			// 反序列化 Redis 中的数据
			err = json.Unmarshal(data, emptyPtr)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal data"})
				return
			}

			// 返回缓存中的数据
			c.JSON(http.StatusOK, gin.H{
				"message": "from redis",
				"data":    emptyPtr,
			})
		}
	}
}
