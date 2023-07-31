package main

import (
	"GrpcStudy/protobuf/validate/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	//var opts []grpc.DialOption
	//opts = append(opts, grpc.WithInsecure())
	dial, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	client := proto.NewGreeterClient(dial)
	ctx, cancelFunc := context.WithTimeout(context.TODO(), 1e9)
	defer cancelFunc()
	rsp, err := client.SayHello(ctx, &proto.Person{
		Id:     1000,
		Email:  "booby@qq.com",
		Mobile: "13385293329",
	})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			if s.Code() == codes.DeadlineExceeded {
				fmt.Println("请求超时")
				return
			}
		}
		panic(err)
	}
	fmt.Println(rsp)
}
