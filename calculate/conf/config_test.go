package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestGetConfigDetail(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	configPath := filepath.Join(path, "conf.yaml")
	fmt.Println(configPath)

	fmt.Println("configpath:", configPath)

	// 获取配置文件内容
	config := viper.New()
	config.SetConfigFile(configPath)
	err = config.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("config:%v", config.AllSettings())
	t.Logf("database host:%v", config.GetString("database.host"))
	t.Logf("database user:%v", config.GetString("database.user"))
	t.Log("database password:", config.Get("database.pwd"))
	t.Log("database dbname:", config.GetString("database.dbname"))

}
