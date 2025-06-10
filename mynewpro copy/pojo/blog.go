package pojo

import (
	"golang_learning/mynewpro/db"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model        //gorm.Model 是一个基础结构体，包含了ID, CreatedAt, UpdatedAt, DeletedAt 等字段
	Id         int    `json:"id" `
	Title      string `json:"title"`
	Content    string `json:"content"`
	Category   string `json:"category"`
}

func FindAllBlogs() []Blog {
	var Blog []Blog
	db.DBConnect.Find(&Blog)
	return Blog
}

func FindByBlogId(BlogId string) Blog {
	var Blog Blog
	db.DBConnect.Where("id = ?", BlogId).First(&Blog)
	return Blog
}

// add blog to database
func AddBlog(blog Blog) Blog {
	db.DBConnect.Create(&blog)
	return blog
}

// delete user
func DeleteBlog(blogId string) Blog {
	blog := Blog{}
	db.DBConnect.Where("id = ?", blogId).Delete(&Blog{})
	return blog
}

// update user
func UpdateBlog(blogId string, blog Blog) Blog {
	db.DBConnect.Model(&blog).Where("id = ?", blogId).Update("blog", blog)
	return blog
}
