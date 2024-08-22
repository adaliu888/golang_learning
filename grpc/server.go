package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
)

func Server() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	s := chat.Server{} // 假设Server是实现了ChatService的服务器结构体
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	log.Println("Chat server started on port 9000")
	defer grpcServer.GracefulStop()
	defer lis.Close()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server on port 9000: %v", err)
	}
}
