package main

import (
	"fmt"
	"reflect"
)

// Reader 接口定义了 Read 方法
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer 接口定义了 Write 方法
type Writer interface {
	Write(p []byte) (n int, err error)
}

// 嵌入 Reader 和 Writer 接口
type ReadWriter interface {
	Reader
	Writer
}

type myReadWriter struct{}

func (m *myReadWriter) Read(p []byte) (n int, err error) {
	// 实现 Read 方法
	return
}
func (m *myReadWriter) Write(p []byte) (n int, err error) {
	// 实现 Write 方法
	return
}

func (m *myReadWriter) ReadWriter(p []byte) (n int, err error) {
	// 实现 ReadWriter 方法
	return
}

func main() {
	rw := &myReadWriter{}
	fmt.Println(rw)

	// 通过反射查询类型
	fmt.Println(reflect.TypeOf(rw))

	// 假设我们有一个实现了 ReadWriter 接口的类型

	// 通过反射查询方法集
	t := reflect.TypeOf(rw)

	fmt.Println("Method set of ReadWriter interface:")
	// 遍历所有方法，打印方法名和参数
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
		fmt.Println(t.Method(i).Type) // t.Method(i)!= t.Method

	}
}
