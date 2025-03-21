package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Config 结构体用于映射 JSON 配置
type Config struct {
	FilePath string `json:"FilePath"`
	FileName string `json:"FileName"`
}

func main() {
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 构建配置文件的完整路径
	configFilePath := filepath.Join(currentDir, "config.json")

	// 读取 JSON 文件
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 解析 JSON 数据
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 打印配置
	fmt.Printf("File Path: %s\n", config.FilePath)
	fmt.Printf("File Name: %s\n", config.FileName)
}
