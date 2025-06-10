package service

import (
	"golang_learning/mynewpro/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//var blogList = []pojo.Blog{}

// get blog

func FindAllBlogs(c *gin.Context) {
	//c.JSON(200, userList)
	blogs := pojo.FindAllBlogs()
	c.JSON(http.StatusOK, blogs)
}

// get user by id
func FindByBlogId(c *gin.Context) {
	blog := pojo.FindByUserId(c.Param("id"))
	if blog.Id == 0 {
		c.JSON(http.StatusNotFound, "error: user not found")
	}
	log.Printf("%+v", blog)
	c.JSON(http.StatusOK, blog)

}

// post user
func AddBlog(c *gin.Context) {
	blog := pojo.Blog{}
	err := c.BindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	//添加到数据库

	//userList = append(userList, user)
	newblog := pojo.AddBlog(blog)

	c.JSON(http.StatusOK, newblog)
	//return
}

// delete user
func DeleteBlog(c *gin.Context) {
	blog := pojo.DeleteBlog(c.Param("id"))
	if blog.Id == 0 {
		c.JSON(http.StatusNotFound, "error")
	}
	c.JSON(http.StatusOK, "deleted successfully")

}
