package main

import (
	"GrpcStudy/protobuf/validate/proto"
	context "context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Service struct {
	proto.UnimplementedGreeterServer
}

func (s *Service) SayHello(ctx context.Context, person *proto.Person) (*proto.Person, error) {
	return &proto.Person{
		Id: 32,
	}, nil
}

type Validator interface {
	Validate() error
}

func main() {
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if r, ok := req.(Validator); ok {
			err = r.Validate()
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	server := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(server, &Service{})
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	server.Serve(listen)
}
