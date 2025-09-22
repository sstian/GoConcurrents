package main

import (
	"context"
	"errors"
	"fmt"
	pb "grpc-protobuf-hello/protobuf"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

// SayHello 业务代码
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("not found token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	fmt.Printf("appId:%s,appKey:%s", appId, appKey)
	fmt.Println()
	if appId != "kuangshen" || appKey != "123123" {
		return nil, errors.New("token is not valid")
	}

	return &pb.HelloResponse{ResponseMsg: "hello" + " " + req.RequestName + " " + appId + " " + appKey}, nil
}

func main() {
	// 1. 开启端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("listen error: %v", err)
		fmt.Println()
	}

	// 2. 创建grpc服务，在grpc服务端中注册
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 3. 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}
