package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Database struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Dbname string `yaml:"dbname"`
	Pwd    string `yaml:"pwd"`
}

func GetConfigDetail() {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configpath := filepath.Join(path, "conf.yaml")

	config := viper.New()

	config.AddConfigPath(configpath) //设置读取的文件路径
	config.SetConfigName("conf")     //设置读取的文件名
	config.SetConfigType("yaml")     //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	//打印文件读取出来的内容:

	fmt.Println(config.Get("database.host"))
	fmt.Println(config.Get("database.user"))
	fmt.Println(config.Get("database.dbname"))
	fmt.Println(config.Get("database.pwd"))

	// 获取并使用 config.yaml 中定义的 key-value
	db := Database{}
	err = config.Unmarshal(&db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Config: ", db)

}
