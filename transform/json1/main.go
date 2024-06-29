package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	//jsonStr := `[{"name":"John","age":20},{"name":"Mary","age":25}]`
	jsonStr := `[{"name":"John","age":15}]`

	//声明一个切片来存储Person类型
	//var persons []Person
	//声明一个map来存储json数据
	//var persons map[string]Person
	//声明一个切片来存储map类型
	//var persons []map[string]interface{}

	//声明一个切片来存储Person类型，并使用json.Unmarshal方法将json数据解析到go变量中
	//err := json.Unmarshal([]byte(jsonStr), &persons)
	//if err!= nil {
	//    fmt.Println("json Unmarshal err:", err)
	//    return
	//}
	//fmt.Println("persons:", persons)

	//声明一个map来存储json数据，并使用json.Unmarshal方法将json数据解析到go变量中
	//err := json.Unmarshal([]byte(jsonStr), &persons)
	var persons []Person
	fmt.Println("persons:", persons)
	//将json数据解析到go变量中
	err := json.Unmarshal([]byte(jsonStr), &persons)

	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}
	fmt.Println("persons:", persons)

	//将go变量转换成json数据
	person := Person{Name: "Jack", Age: 23}
	ls := reflect.TypeOf(person)
	fmt.Print(ls.Name())
	fmt.Println("person:", person)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("json Marshal err:", err)
	}
	fmt.Println("jsonData:", string(jsonData))

	//读取json数据
	file, err := os.Open("person.json")
	if err != nil {
		fmt.Println("create file err:", err)
		return
	}
	defer file.Close()
	jsonData, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read file err:", err)
	}
	fmt.Println("jsonData:", string(jsonData))

	//将json数据写入到文件
	file, err = os.Create("person.json")
	if err != nil {
		fmt.Println("create file err:", err)
		return
	}
	defer file.Close()

	//将go变量转换成json数据并写入到文件
	jsonData, err = json.Marshal(person)
	if err != nil {
		fmt.Println("json Marshal err:", err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("write file err:", err)
	}

	json.NewEncoder(file).Encode(persons)
	fmt.Println("write file success")
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
