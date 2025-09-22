package main

import (
	"context"
	"fmt"
	pb "grpc-protobuf-hello/protobuf"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// TLS认证
	cred, err := credentials.NewClientTLSFromFile("./key/test.pem",
		"*.kuangstudy.com")
	if err != nil {
		fmt.Println(err)
		fmt.Println("NewClientTLSFromFile error")
	}

	// 1. 连接到server端
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(cred)) // 使用安全传输
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
