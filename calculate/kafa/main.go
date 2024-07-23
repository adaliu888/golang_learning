package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
	_ "gopkg.in/yaml.v2"
)

type KafkaCluster struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml: "kind"`
	Metadata   Metadata `yaml: "metadata"`
	Spec       Spec     `yaml: "spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
	//map类型
	Labels map[string]*NodeServer `yaml:"labels"`
}

type NodeServer struct {
	Address string `yaml: "address"`
	Id      string `yaml: "id"`
	Name    string `yaml: "name"`
	//注意，属性里，如果有大写的话，tag里不能存在空格
	//如yaml: "nodeName" 格式是错误的，中间多了一个空格，不能识别的
	NodeName string `yaml:"nodeName"`
	Role     string `yaml: "role"`
}

type Spec struct {
	Replicas int    `yaml: "replicas"`
	Name     string `yaml: "name"`
	Image    string `yaml: "iamge"`
	Ports    int    `yaml: "ports"`
	//slice类型
	Conditions []Conditions `yaml: "conditions"`
}

type Conditions struct {
	ContainerPort string   `yaml:"containerPort"`
	Requests      Requests `yaml: "requests"`
	Limits        Limits   `yaml: "limits"`
}

type Requests struct {
	CPU    string `yaml: "cpu"`
	MEMORY string `yaml: "memory"`
}

type Limits struct {
	CPU    string `yaml: "cpu"`
	MEMORY string `yaml: "memory"`
}

// 主程序入口
func main() {
	var c KafkaCluster
	//读取yaml配置文件, 将yaml配置文件，转换struct类型
	conf := c.getConf()

	//将对象，转换成json格式
	data, err := json.Marshal(conf)

	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}

	//最终以json格式，输出
	fmt.Println("data:\t", string(data))
}

// 读取Yaml配置文件，并转换成KafkaCluster对象  struct结构
func (kafkaCluster *KafkaCluster) getConf() *KafkaCluster {

	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	vip := viper.New()
	vip.AddConfigPath(path + "/config") //设置读取的文件路径
	vip.SetConfigName("application")    //设置读取的文件名
	vip.SetConfigType("yaml")           //设置文件的类型
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}

	err = vip.Unmarshal(&kafkaCluster)
	if err != nil {
		panic(err)
	}

	return kafkaCluster
}
