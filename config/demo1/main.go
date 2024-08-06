package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	// 使用相对路径
	relativePath := "config/config.yaml"
	// 使用 filepath 包来确保路径在 Windows 下正确
	fullPath, _ := filepath.Abs(relativePath)
	fmt.Println("Absolute path:", fullPath)

	// 使用绝对路径
	absolutePath := filepath.Join(
		"E:",
		`\GitHub\golang_learning\src\golang_learning\`,
		"config",
		"config.yaml",
	)
	fmt.Println("Absolute path:", absolutePath)
	content, err := ioutil.ReadFile(absolutePath)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(string(content))

	type DBconfig struct {
		Host      string `yaml:"host" json:"host"`
		port      int    `yaml:"port" 	json:"port"`
		usernames string `yaml:"usernames" json:"usernames"`
		password  string `yaml:"password" json:"password"`
		dbname    string `"yaml:"dbname" json:"dbname"`
	}

	var dbConfig DBconfig
	err = yaml.Unmarshal(content, &dbConfig)
	if err != nil {
		log.Println(err)
	}
	log.Printf("DBConfig: %+v\n", dbConfig)

}
