package main

import (
	"GrpcStudy/protobuf/validate/proto"
	context "context"
	"fmt"
	"github.com/xcltapestry/golibs/utils/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"time"
)

type Service struct {
	proto.UnimplementedGreeterServer
}

func mockServiceWithTimeOut(ctx context.Context) {
	// 执行service
	fmt.Println("执行service logic")
}

func (s *Service) SayHello(ctx context.Context, person *proto.Person) (*proto.Person, error) {
	time.Sleep(0.5e9)
	fmt.Println("先执行主体业务逻辑")

	deadline := ctxutil.ShrinkDeadline(ctx, time.Millisecond*200)
	sub := deadline.Sub(time.Now())
	fmt.Println(sub)
	ctxNew, cancelFunc := context.WithTimeout(context.TODO(), sub)
	defer cancelFunc()

	mockServiceWithTimeOut(ctxNew)

	return &proto.Person{
		Id: 32,
	}, status.Errorf(codes.NotFound, "not found person")
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
