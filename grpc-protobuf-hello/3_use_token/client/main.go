package main

import (
	"context"
	"fmt"
	pb "grpc-protobuf-hello/protobuf"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
// C:\Users\Snow-Angel\go\pkg\mod\google.golang.org\grpc@v1.75.1\credentials\credentials.go

// PerRPCCredentials defines the common interface for the credentials which need to
// attach security information to every RPC (e.g., oauth2).
type PerRPCCredentials interface {

	// GetRequestMetadata gets the current request metadata, refreshing tokens
	// if required. This should be called by the transport layer on each
	// request, and the data should be populated in headers or other
	// context. If a status code is returned, it will be used as the status for
	// the RPC (restricted to an allowable set of codes as defined by gRFC
	// A54). uri is the URI of the entry point for the request.  When supported
	// by the underlying implementation, ctx can be used for timeout and
	// cancellation. Additionally, RequestInfo data will be available via ctx
	// to this call.  TODO(zhaoq): Define the set of the qualified keys instead
	// of leaving it as an arbitrary string.
	// 第一个方法作用是获取元数据信息，也就是客户端提供的key,value对。context用于控制超时和取消，uri是请求入口处的uri。
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)

	// RequireTransportSecurity indicates whether the credentials requires
	// transport security.
	// 第二个方法的作用是否需要基于 TLS 认证进行安全传输。如果返回值是true，则必须加上TLS验证;返回值是false则不用。
	RequireTransportSecurity() bool
}
*/

// ClientTokenAuth 自定义token
type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "kuangshen",
		"appKey": "123123",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	var dials []grpc.DialOption
	dials = append(dials, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dials = append(dials, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	// 1. 连接到server端
	conn, err := grpc.Dial("127.0.0.1:9090", dials...) // 使用token
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 2. 建立连接
	client := pb.NewSayHelloClient(conn)

	// 3. 执行rpc调用
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "jack"})
	fmt.Printf(resp.ResponseMsg) //hello jack kuangshen 123123
}
