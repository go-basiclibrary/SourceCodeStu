package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"gobasic/protobuf/grpc_resolver/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(context.Context, *proto.Person) (*proto.Person, error) {
	fmt.Println("我进来了,hello,hello")
	return &proto.Person{
		Id:     1,
		Email:  "821123693@qq.com",
		Mobile: "15044785521",
	}, nil
}

// grpc 服务端,注册服务 to consul
func main() {
	lis, err := net.Listen("tcp", ":7890")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	// 注册服务到consul上
	registerServerConsul()

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

// consul服务注册
func registerServerConsul() {
	// 创建consul客户端
	cfg := api.DefaultConfig()
	cfg.Address = "43.143.172.37:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	registration := new(api.AgentServiceRegistration)
	registration.ID = "test-consul-for-trpc"
	registration.Name = "Test"
	registration.Port = 7890
	registration.Tags = []string{"test", "consul", "grpc", "解析器"}
	registration.Address = "127.0.0.1"

	// 注册服务
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}
