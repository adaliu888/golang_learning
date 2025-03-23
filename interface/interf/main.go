package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type myReadWriter struct{}

func (m *myReadWriter) Read(p []byte) (n int, err error) {
	f, err := os.Open("test.txt")
	if err != nil { //错误处理
		return
	}
	defer f.Close() //延迟关闭
	return
}

func (m *myReadWriter) Write(p []byte) (n int, err error) {
	f, err := os.Create("test.txt") //
	if err != nil {
		fmt.Println("the file is exist alrea")
	}
	defer f.Close()
	return
}

func main() {
	rw := &myReadWriter{}
	rw.Write([]byte("test.txt"))
	rw.Read([]byte("test.txt"))

	fmt.Println(rw)
}
