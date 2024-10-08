package main

import (
	"golang_learning/mynewpro/db"
	"golang_learning/mynewpro/file"
	"golang_learning/mynewpro/middlewave"
	src "golang_learning/mynewpro/src_router"

	"github.com/gin-gonic/gin"
)

/*
	func setlogger() {
		f, _ := os.Create("./gin.log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
*/
func main() {
	//setlogger()

	//ZapLogger := middlewave.NewLogger()
	//ZapLogger.Info("this is an info level log")
	//db.DBIint()
	//启动数据库

	Log := middlewave.InitLogger(file.FN(), "info")
	Log.Info(":info,server started")

	//setlogger()
	db.DBIint()
	db.InitRedis()
	Log.Info("db initialized")

	router := gin.Default() //创建路由

	//add swagger
	//router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//router.Use(middlewave.Logger())
	//router.Use(middlewave.GinLogger(lg), middlewave.GinRecovery(lg, true))
	//user router
	router.Use(middlewave.RateLimitMiddleware()) //访问限流
	v1 := router.Group("/v1")                    //分组
	src.AddUserRouter(*v1)
	//blog router
	v2 := router.Group("/v2") //
	src.AddBlogRouter(*v2)

	//router.Use(middlewave.ZapLogger()) //添加路由
	//router.Use(gin.BasicAuth(gin.Accounts{"admin": "admin"}), middlewave.Logging())
	//中间件logger,需要登录才能访问
	router.Use(gin.BasicAuth(gin.Accounts{"admin": "admin"}))
	//启动数据库
	go func() {
		db.DBIint()
	}()

	router.Run() // listen and serve on 0.0.0.0:8080

}
