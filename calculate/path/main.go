package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// database struct
type DBConfig struct {
	Host string
	Port int
}

// 使用viper读取config.yaml
func InitMysqlConfig(env string) DBConfig {
	//设置配置文件的路径（可以是相对路经，也可以是绝对路径）
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/dev")
	//设置配置文件名称
	viper.SetConfigName("config")
	//设置配置文件类型
	viper.SetConfigType("yaml") // 根据环境变量选择加载不同的配置文件

	if env == "dev" {
		viper.Set("env", "dev") // 可以设置一个标记，以便在配置文件中使用
	}

	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error")

	}

	// 从配置文件中读取数据，并转换为正确的类型
	host := viper.GetString("mysql.host")
	if host == "" {
		log.Fatal("MySQL host is not set in the config file")
	}
	port := viper.GetInt("mysql.port")
	if port == 0 {
		log.Fatal("MySQL port is not set in the config file")
	}

	return DBConfig{
		Host: host,
		Port: port,
	}
}

// 获取文件的基本信息from GetFileInfo()
type FileDetail struct {
	Name    string
	Size    int64
	Mode    fs.FileMode
	ModTime time.Time
	Isdir   bool
}

func main() {
	//指定文件路径
	path := "./example.txt"
	//获取文件信息
	FileDetail := GetFileInfo(path)
	fmt.Println(FileDetail)
	//配置环境
	env := "dev"
	mysqlconfig := InitMysqlConfig(env)
	fmt.Println(mysqlconfig.Host, mysqlconfig.Port)

}

func GetFileInfo(filepath string) FileDetail {

	FileInfo, err := os.Stat(filepath)
	fmt.Println(FileInfo)
	if err != nil {
		log.Println("Error:", err)
	}
	return FileDetail{
		Name:    FileInfo.Name(),
		Size:    FileInfo.Size(),
		Mode:    FileInfo.Mode(),
		ModTime: FileInfo.ModTime(),
		Isdir:   FileInfo.IsDir(),
	}

}
