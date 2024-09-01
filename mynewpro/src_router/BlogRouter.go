package src_router

//集中管理路由
import (
	"golang_learning/mynewpro/service"

	"github.com/gin-gonic/gin"
)

func AddBlogRouter(r gin.RouterGroup) {
	blog := r.Group("/blogs")
	blog.GET("/", service.FindAllBlogs)
	blog.GET("/:id", service.FindByBlogId)
	blog.POST("/register", service.AddBlog)
	blog.DELETE("/:id", service.DeleteBlog)
	/* blog.GET("/:id/comments", FindBlogComments)
	   blog.POST("/:id/comments", CreateBlogComment)
	   blog.DELETE("/:id/comments/:commentId", DeleteBlogComment)
	   blog.GET("/:id/tags", FindBlogTags)
	   blog.POST("/:id/tags", CreateBlogTag)
	   blog.DELETE("/:id/tags/:tagId", DeleteBlogTag)
	   blog.GET("/:id/likes", FindBlogLikes)
	   blog.POST("/:id/likes", LikeBlog)

	*/
}
