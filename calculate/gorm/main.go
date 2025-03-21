package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 定义一个模型
type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	// 连接到 SQLite 数据库
	db, err := gorm.Open(sqlite.Open("go-test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 检查表是否存在
	if !db.Migrator().HasTable(&User{}) {
		// 如果表不存在，则迁移模式
		db.AutoMigrate(&User{})
	}

	// 创建记录
	user := User{Name: "Alice", Age: 30}
	db.Create(&user)

	// 查询记录
	var users []User
	db.Find(&users)
	log.Println(users)
}
