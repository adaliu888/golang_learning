package conf

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/viper"
)

// DatabaseConfig 包含数据库配置
type DatabaseConfig struct {
    Host   string `yaml:"host"`
    User   string `yaml:"user"`
    DBName string `yaml:"dbname"`
    Pwd    string `yaml:"pwd"`
}

// Config 包含所有应用配置
type Config struct {
    Database DatabaseConfig `yaml:"database"`
    // 可以添加其他配置部分，如JWT等
}

// LoadConfig 从配置文件加载配置
func LoadConfig(configPath string) (*Config, error) {
    // 如果没有提供配置路径，使用默认路径
    if configPath == "" {
        path, err := os.Getwd()
        if err != nil {
            return nil, fmt.Errorf("获取工作目录失败: %w", err)
        }
        configPath = filepath.Join(path, "config.yaml")
    }

    config := viper.New()
    config.SetConfigFile(configPath)
    config.SetConfigType("yaml")

    // 尝试读取配置
    if err := config.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("读取配置文件失败: %w", err)
    }

    // 解析配置到结构体
    var appConfig Config
    if err := config.Unmarshal(&appConfig); err != nil {
        return nil, fmt.Errorf("解析配置失败: %w", err)
    }

    fmt.Printf("数据库配置: %s, %s, %s, %s\n", 
        appConfig.Database.Host, 
        appConfig.Database.User, 
        appConfig.Database.DBName, 
        appConfig.Database.Pwd)

    return &appConfig, nil
}

