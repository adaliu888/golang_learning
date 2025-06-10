package main

import (
	"golang_learning/mynewpro/db"
	"golang_learning/mynewpro/file"
	"golang_learning/mynewpro/jwt"
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

	// 初始化 JWT 配置
	if err := jwt.InitJWTConfig(); err != nil {
		Log.Error("Failed to initialize JWT config: " + err.Error())
	}

	//setlogger()
	db.DBIint()
	db.InitRedis()
	Log.Info("db initialized")

	router := gin.Default() //创建路由

	//add swagger
	//router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//router.Use(middlewave.Logger())
	//router.Use(middlewave.GinLogger(lg), middlewave.GinRecovery(lg, true))

	// 为API应用速率限制中间件
	router.Use(middlewave.RateLimitMiddleware()) //访问限流

	// 用户相关API
	v1 := router.Group("/v1")
	src.AddUserRouter(*v1)

	// 博客相关API
	v2 := router.Group("/v2")
	src.AddBlogRouter(*v2)

	// 认证相关API (JWT模式)
	auth := router.Group("/auth")
	src.AddAuthRouter(*auth)

	// 限制访问管理员界面
	admin := router.Group("/admin")
	admin.Use(gin.BasicAuth(gin.Accounts{"admin": "admin"}))
	{
		// 管理员路由...
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}
