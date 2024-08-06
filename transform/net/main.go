package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// 连接到远程服务器
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 设置超时时间
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	// 发送请求
	request := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		panic(err)
	}

	// 读取响应
	resp := make([]byte, 4096)
	n, err := conn.Read(resp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received %d bytes: %s\n", n, resp[:n])

	file := "test.txt"
	err = os.WriteFile(file, resp[:n], 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saved response to %s\n", file)
}
