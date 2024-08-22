package grpc

import (
	"log"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// 创建 gRPC 客户端
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	// 构建 gRPC 请求
	message := chat.Message{Body: "Hello from client"}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 处理 gRPC 响应
	log.Printf("Greeting: %s", response.Body)

}
