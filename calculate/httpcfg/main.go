package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type HttpClientConfig struct {
	Enable         bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	EnableAutoTest bool   `mapstructure:"enable_autotest" json:"enable_autotest" yaml:"enable_autotest"`
	EnableDebug    bool   `mapstructure:"enable_debug" json:"enable_debug" yaml:"enable_debug"`
	Name           string `mapstructure:"name" json:"name" yaml:"name"`
	Tag            string `mapstructure:"tag" json:"tag" yaml:"tag"`

	Url                string `mapstructure:"url" json:"url" yaml:"url"`
	ConnectReadTimeout int    `mapstructure:"connect-read-timeout" json:"connect-read-timeout" yaml:"connect-read-timeout"`
	DoTimeInterval     int    `mapstructure:"do-timeinterval" json:"do-timeinterval" yaml:"do-timeinterval"`
}

func main() {
	// 假设我们有一个名为 config.json 的配置文件
	configFile := "config.json"

	// 读取配置文件
	fileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	// 解码 JSON 到结构体
	var config HttpClientConfig
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		panic(err)
	}

	// 使用配置
	fmt.Printf("%+v\n", config)
}
