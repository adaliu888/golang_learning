package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// 设置配置文件的路径和名称
	viper.SetConfigFile("./config.yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 获取配置的值
	appPort := viper.Get("app.port")
	dbName := viper.Get("database.dbname")

	// 使用配置信息
	fmt.Printf("Application Port: %v\n", appPort)
	fmt.Printf("Database Name: %v\n", dbName)
}
