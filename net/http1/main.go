package main

import (
	"fmt"
	"net"
)

func connectToServer(address string, index int) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Error connecting: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Printf("Connection %d established\n", index)
	// 执行一些操作
}

func main() {
	address := "127.0.0.1:8080"
	for i := 0; i < 10; i++ {
		go connectToServer(address, i) // 使用 Goroutine 创建连接
	}
}
