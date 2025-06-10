package src_router

//集中管理路由
import (
	md "golang_learning/mynewpro/middlewave"
	"golang_learning/mynewpro/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r gin.RouterGroup) {
	user := r.Group("/users")

	// 公共访问端点
	user.GET("/", service.FindAllUsers)
	user.GET("/:id", service.FindByUserId)
	user.POST("/register", service.PostUser)
	user.POST("/login", service.Login)
	user.GET("/logout", service.Logout)

	// JWT刷新令牌端点
	user.POST("/refresh-token", md.RefreshTokenHandler)

	// 兼容旧的会话验证的路由
	sessionProtected := user.Group("/", md.SetSession())
	sessionProtected.GET("/check", service.CheckUserSession)

	sessionProtected.Use(md.AuthSession())
	{
		sessionProtected.DELETE("/:id", service.DeleteUser)
		sessionProtected.PUT("/:id", service.UpdateUser)
	}

	// JWT保护的路由 - 可以根据需要添加
	jwtProtected := user.Group("/jwt")
	jwtProtected.Use(md.AuthTokenMiddleware())
	{
		jwtProtected.GET("/profile", service.FindAllUsers) // 示例，可以替换为具体的JWT保护资源
		jwtProtected.DELETE("/:id", service.DeleteUser)
		jwtProtected.PUT("/:id", service.UpdateUser)
	}
}

// 添加用于演示的处理函数
func AddAuthRouter(r gin.RouterGroup) {
	auth := r.Group("/auth")

	// 公共访问端点
	auth.POST("/login", service.Login)
	auth.POST("/refresh-token", md.RefreshTokenHandler)
	auth.GET("/logout", md.LogoutHandler)

	// JWT保护的资源
	protected := auth.Group("/protected")
	protected.Use(md.AuthTokenMiddleware())
	{
		protected.GET("/resource", func(c *gin.Context) {
			userID := c.GetString("userID")
			c.JSON(200, gin.H{
				"message": "You accessed a protected resource",
				"userID":  userID,
			})
		})
	}
}

//func AddBlogRouter(r gin.RouterGroup) {
//	blog := r.Group("/blogs", md.AuthMiddleware())
//	blog.GET("/", service.FindAllBlogs)
//	blog.GET("/:id", service.FindByBlogId)
//}
