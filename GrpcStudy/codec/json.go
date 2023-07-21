package main

import (
	"GrpcStudy/codec/proto"
	"context"
	"encoding/json"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type JSONCodec struct{}

func (c *JSONCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *JSONCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (c *JSONCodec) Name() string {
	return "json"
}

func main() {
	// 创建一个gRPC服务器
	server := grpc.NewServer(
		grpc.ForceServerCodec(&JSONCodec{}), // 设置自定义编解码器为JSON
	)

	proto.RegisterGreeterServer(server, &MyService{})

	// 启动监听
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	// 开始接受请求
	if err := server.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return
	}
}

// MyService 实现gRPC服务的具体逻辑
type MyService struct {
	proto.UnimplementedGreeterServer
}

func (s *MyService) SayHello(ctx context.Context, person *proto.Person) (*proto.Person, error) {
	return &proto.Person{
		Id:     1,
		Email:  person.Email,
		Mobile: person.Mobile,
	}, nil
}

//func (s *MyService) mustEmbedUnimplementedGreeterServer() {
//	//TODO implement me
//	panic("implement me")
//}

//func (s *MyService) MyMethod(ctx context.Context, req *example.MyMessage) (*example.MyMessage, error) {
//	fmt.Printf("Received request: %v\n", req)
//	return &example.MyMessage{Data: "Hello, gRPC!"}, nil
//}
