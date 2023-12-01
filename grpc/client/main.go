package main

import (
	pb "MyGoProject/grpc/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "lijianhua",
		"appKey": "123123",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("127.0.0.1:8088", opts...)
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	// 执行rpc调用（这个方法在服务器端来实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "lijianhua"})

	fmt.Println(resp.GetResponseMsg())
}
