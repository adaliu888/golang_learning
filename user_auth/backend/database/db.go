package database

import (
	"fmt"
	"golang_learning/user_auth/backend/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB 是全局数据库连接
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	var err error
	dbType := os.Getenv("DB_TYPE")
	dbName := os.Getenv("DB_NAME")

	if dbType == "sqlite" {
		DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	} else {
		return fmt.Errorf("unsupported database type: %s", dbType)
	}

	if err != nil {
		return err
	}

	// 自动迁移
	if err := DB.AutoMigrate(&models.User{}, &models.Doctor{}); err != nil {
		return err
	}

	return nil

}
