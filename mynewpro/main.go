package main

import (
	"golang_learning/mynewpro/db"
	"golang_learning/mynewpro/middlewave"
	src "golang_learning/mynewpro/src_router"

	"github.com/gin-gonic/gin"
)

/*func setlogger() {
	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
*/

func main() {

	//ZapLogger := middlewave.NewLogger()
	//ZapLogger.Info("this is an info level log")
	db.DBIint()
	//启动数据库
	lg := middlewave.InitLogger()
	//setlogger()

	router := gin.Default() //创建路由
	//router.Use(middlewave.Logger())
	router.Use(middlewave.GinLogger(lg), middlewave.GinRecovery(lg, true))
	router.Use(middlewave.RateLimitMiddleware()) //访问限流
	v1 := router.Group("/v1")                    //分组
	src.AddUserRouter(*v1)

	//v2 := router.Group("/v2")
	//src.AddBlogRouter(*v2)
	//router.Use(middlewave.ZapLogger()) //添加路由
	//router.Use(gin.BasicAuth(gin.Accounts{"admin": "admin"}), middlewave.Logger())
	//中间件logger,需要登录才能访问
	router.Use(gin.BasicAuth(gin.Accounts{"admin": "admin"}), middlewave.Logger())
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
