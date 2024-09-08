package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Database struct {
	host   string `yaml:"host"`
	user   string `yaml:"user"`
	dbname string `yaml:"dbname"`
	pwd    string `yaml:"pwd"`
}

func ConfigYaml() {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()

	config.AddConfigPath(path)     //设置读取的文件路径
	config.SetConfigName("config") //设置读取的文件名
	config.SetConfigType("yaml")   //设置文件的类型

	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	// 读取yaml中的数据
	db := Database{}
	err = config.Unmarshal(&db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Config: ", db)
}
