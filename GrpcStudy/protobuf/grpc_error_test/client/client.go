package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"gobasic/protobuf/grpc_error_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func main() {
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithMax(3),                                                          // 最大重试次数
		grpc_retry.WithPerRetryTimeout(1 * time.Second),                                // 超时时间
		grpc_retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Unavailable), // 什么状态下重试
	}

	//var opts []grpc.DialOption
	//opts = append(opts, grpc.WithInsecure())
	var opts []grpc.DialOption
	//这个请求应该多长时间超时,这个重试应该几次,当服务器返回什么状态码的时候重试
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)), grpc.WithInsecure())
	dial, err := grpc.Dial("127.0.0.1:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	client := proto.NewGreeterClient(dial)
	//ctx, _ := context.WithTimeout(context.TODO(), 1e9)

	ctx := context.TODO()
	rsp, err := client.SayHello(ctx, &proto.Person{
		Id:     1000,
		Email:  "booby@qq.com",
		Mobile: "15022331145",
	},
	)
	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			panic("解析失败")
		}
		fmt.Println(status.Message())
		fmt.Println(status.Code())
	}
	fmt.Println(rsp)
	fmt.Println(ctx.Err())
	time.Sleep(3e9)
}
