package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

func main() {
	md := metadata.New(map[string]string{
		"opentracing": "letabc",
		"sName":       "ws",
	})

	ctx, cancelFunc := context.WithCancel(context.TODO())
	defer cancelFunc()

	ctx = metadata.NewIncomingContext(ctx, md)
	md, b := metadata.FromIncomingContext(ctx)
	if b {
		fmt.Println(b, md)
	}
	fmt.Println(md)
}
