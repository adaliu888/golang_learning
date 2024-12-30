package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `json:"id" `
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	// 连接到MySQL数据库
	// 请替换以下参数为你的数据库信息
	dsn := "root:test1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// 自动迁移模式，GORM将会自动创建User表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("failed to migrate:", err)
	}

	fmt.Println("User table created successfully")
}
