package main

import (
	"context"
	"fmt"
	pb "grpc-protobuf-hello/protobuf"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello" + " " + req.RequestName}, nil
}

func main() {
	// 1. 开启端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("listen error: %v", err)
		fmt.Println()
	}

	// 2. 创建grpc服务，在grpc服务端中注册
	grpcServer := grpc.NewServer()
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 3. 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}
