package main

//json 转换 解析 写入 写出
import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 将json 数据解析到go变量中去
func main() {
	jsonStr := `[{"name":"John","age":20},{"name":"Mary","age":25}]`
	var persons []Person
	//将json数据解析到go变量中
	err := json.Unmarshal([]byte(jsonStr), &persons)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
	}
	fmt.Println("persons:", persons)

	//将go变量转换成json数据
	person := Person{Name: "Jack", Age: 23}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("json Marshal err:", err)
	}
	fmt.Println("jsonData:", string(jsonData))

	//将go变量转换成json数据并写入到文件
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println("create file err:", err)
		return
	}
	defer file.Close()
	jsonData, err = json.Marshal(person)
	if err != nil {
		fmt.Println("json Marshal err:", err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("write file err:", err)
	}
}
