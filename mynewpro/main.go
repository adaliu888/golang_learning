package main

import (
	"io"
	"mynewpro/db"
	"mynewpro/middlewave"
	"mynewpro/src"
	"os"

	"github.com/gin-gonic/gin"
)

func setlogger() {
	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setlogger() //设置日志
	//启动数据库
	go func() {
		db.DBIint()
	}()

	router := gin.Default()                      //创建路由
	router.Use(middlewave.RateLimitMiddleware()) //访问限流
	v1 := router.Group("/v1")                    //分组
	src.AddUserRouter(*v1)

	//v2 := router.Group("/v2")
	//src.AddBlogRouter(*v2)
	//router.Use(middlewave.ZapLogger())                                             //添加路由
	router.Use(gin.BasicAuth(gin.Accounts{"admin": "admin"}), middlewave.Logger()) //中间件logger,需要登录才能访问

	//启动数据库
	go func() {
		db.DBIint()
	}()

	/*
		r.GET("/ping", func(c *gin.Context) { //get ping message
			c.JSON(200, gin.H{
				"message": "ping",
				"version": "1.0.0",
			})
		})
		r.POST("/ping/:id", func(c *gin.Context) { //post ping message
			id := c.Param("id")
			c.JSON(200, gin.H{
				"id": id,
			})

		})*/

	router.Run() // listen and serve on 0.0.0.0:8080
}
