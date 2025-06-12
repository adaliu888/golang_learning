package conf

import (
	"fmt"
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
	fmt.Printf("当前工作目录: %s\n", path)

	// 直接指定完整的配置文件路径
	configFile := filepath.Join(path, "config.yaml")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Fatalf("配置文件不存在: %s", configFile)
	}

	// 设置配置文件的路径和名称
	v := viper.New()
	v.SetConfigFile(configFile) // 直接设置配置文件的完整路径

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			t.Fatalf("配置文件未找到: %v\n使用的配置路径: %s", err, path)
		} else {
			t.Fatalf("读取配置文件错误: %v", err)
		}
	}

	fmt.Printf("使用的配置文件: %s\n", v.ConfigFileUsed())

	// 获取所有配置，用于调试
	allSettings := v.AllSettings()
	fmt.Printf("所有配置: %+v\n", allSettings)

	// 获取并使用 config.yaml 中定义的 key-value
	host := v.GetString("database.host")
	user := v.GetString("database.user")
	dbname := v.GetString("database.dbname")
	pwd := v.GetString("database.pwd")

	// 验证配置值是否正确读取
	if host != "127.0.0.1" {
		t.Errorf("期望 database.host 为 '127.0.0.1'，实际为 '%s'", host)
	}
	if user != "root" {
		t.Errorf("期望 database.user 为 'root'，实际为 '%s'", user)
	}

	fmt.Printf("数据库配置: host=%s, user=%s, dbname=%s, pwd=%s\n", 
		host, user, dbname, pwd)
}