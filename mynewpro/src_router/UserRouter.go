package src

//集中管理路由
import (
	md "golang_learning/mynewpro/middlewave"
	"golang_learning/mynewpro/pojo"
	"golang_learning/mynewpro/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r gin.RouterGroup) {
	user := r.Group("/users", md.SetSession())

	user.GET("/", service.FindAllUsers)
	//user.GET("/:id", service.FindByUserId)
	user.GET("/:id", service.CacheOneUserDecorator(service.RedisOneUser, "id", "user_%s", pojo.User{}))
	user.POST("/register", service.PostUser)
	//delete user
	//user.DELETE("/:id", service.DeleteUser)
	//update user

	//user.PUT("/:id", service.UpdateUser)

	//login
	user.POST("/login", service.Login)

	user.GET("/logout", service.Logout)

	user.GET("/check", service.CheckUserSession)

	user.Use(md.AuthSession())
	{
		//delete user
		user.DELETE("/:id", service.DeleteUser)
		//update user
		user.PUT("/:id", service.UpdateUser)

	}

}

//func AddBlogRouter(r gin.RouterGroup) {
//	blog := r.Group("/blogs", md.AuthMiddleware())
//	blog.GET("/", service.FindAllBlogs)
//	blog.GET("/:id", service.FindByBlogId)
//}
