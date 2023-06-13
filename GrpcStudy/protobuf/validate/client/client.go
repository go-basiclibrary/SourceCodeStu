package main

import (
	"GrpcStudy/protobuf/validate/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	rsp, err := client.SayHello(context.Background(), &proto.Person{
		Id:     1000,
		Email:  "booby@qq.com",
		Mobile: "13385293329",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}
