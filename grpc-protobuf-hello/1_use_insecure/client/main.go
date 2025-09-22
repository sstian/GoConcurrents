package main

import (
	"context"
	"fmt"
	pb "grpc-protobuf-hello/protobuf"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 连接到server端
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials())) // 此处禁用安全传输，没有加密和验证
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 2. 建立连接
	client := pb.NewSayHelloClient(conn)

	// 3. 执行rpc调用
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "jack"})
	fmt.Printf(resp.ResponseMsg) //hello jack
}
