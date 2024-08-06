/*
在Go语言中，实现UDP协议的多播通信涉及以下几个步骤：

创建UDP套接字：
使用net.ListenUDP或net.DialUDP创建一个UDP套接字。

加入多播组：
使用setsockopt设置套接字选项，加入多播组。

绑定多播地址：
使用ListenUDP的地址和端口进行绑定。

接收多播数据：
使用ReadFromUDP或Read方法接收多播数据。

发送多播数据：
使用WriteToUDP或Write方法发送多播数据。

退出多播组：
使用setsockopt设置套接字选项，退出多播组。

关闭套接字：
使用Close方法关闭UDP套接字。

以下是实现UDP多播通信的示例代码：
*/
package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 多播地址和端口
	addr := "224.0.0.1:8080"

	// 创建UDP套接字
	conn, err := net.ListenMulticastUDP("udp", nil, &net.UDPAddr{Port: 8080})
	if err != nil {
		println("Failed to create UDP connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 设置多播选项
	if err := conn.JoinGroup(nil, &net.UDPAddr{IP: net.ParseIP("224.0.0.1")}); err != nil {
		println("Failed to join group:", err)
		os.Exit(1)
	}
	defer conn.LeaveGroup(nil, &net.UDPAddr{IP: net.ParseIP("224.0.0.1")})

	// 监听中断信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		println("Received interrupt, exiting...")
		conn.LeaveGroup(nil, &net.UDPAddr{IP: net.ParseIP("224.0.0.1")})
		os.Exit(0)
	}()

	// 接收多播数据
	var buf [512]byte
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			println("Failed to read from UDP:", err)
			continue
		}
		println("Received message from", remoteAddr, ":", string(buf[:n]))
	}
}

/*在这个示例中，我们首先创建了一个UDP套接字，并使用ListenMulticastUDP指定多播地址和端口。然后，我们使用JoinGroup方法加入多播组，并使用LeaveGroup方法在退出时离开多播组。

我们监听中断信号，以便在接收到信号时退出多播组并关闭程序。

在接收循环中，我们使用ReadFromUDP方法接收多播数据，并打印发送者的地址和接收到的消息。

请注意，多播通信可能受到网络环境和操作系统的限制，因此在实际应用中可能需要进行适当的配置和测试。*/
