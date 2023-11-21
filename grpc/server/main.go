package main

import (
	pb "MyGoProject/grpc/server/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	fmt.Println(appId, appKey)
	if appId != "lijianhua" || appKey != "123" {
		return nil, errors.New("token 不正确")
	}

	return &pb.HelloResponse{ResponseMsg: "hello," + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":8088")
	// 创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	// 在grpc服务端中注册我们自己写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		return
	}

}
