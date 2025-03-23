package main

// go json.Marshal(t) 是 Go 语言 encoding/json 包中的一个函数，用于将 Go 中的数据类型转换成 JSON 格式的字节序列
//指针和空、匿名字段，不会被打印序列
import (
	"encoding/json"
	"fmt"
)

func main() {
	type T struct {
		F1 int `json:"f_1"`
		F2 int `json:"f_2,omitempty"`
		F3 int `json:"f_3,omitempty"`
		F4 int `json:"-"`
	}
	t := T{1, 0, 2, 3}
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b) // {"f_1":1,"f_3":2}
}
