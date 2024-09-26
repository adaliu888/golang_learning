package conf_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestConfigYaml(t *testing.T) {
	// 获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	path = filepath.Join(path, "config.yaml")

	// 设置配置文件的路径和名称
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found; using defaults")
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	// 获取并使用 config.yaml 中定义的 key-value
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")

	fmt.Printf("Database.host: %v, database.port: %v\n", host, port)
}
