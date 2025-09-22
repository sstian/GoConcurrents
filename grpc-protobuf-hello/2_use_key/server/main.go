package main

import (
	"context"
	"fmt"
	pb "grpc-protobuf-hello/protobuf"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello" + " " + req.RequestName}, nil
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Println("Current working directory:", cwd)

	// TLS认证
	cred, err := credentials.NewServerTLSFromFile("./key/test.pem", "./key/test.key")
	if err != nil {
		fmt.Println(err)
		fmt.Println("NewServerTLSFromFile error")
	}

	// 1. 开启端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("listen error: %v", err)
		fmt.Println()
	}

	// 2. 创建grpc服务，在grpc服务端中注册
	grpcServer := grpc.NewServer(grpc.Creds(cred))
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 3. 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}
