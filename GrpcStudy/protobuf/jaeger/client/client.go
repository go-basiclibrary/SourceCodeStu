package main

import (
	"GrpcStudy/protobuf/jaeger/otgrpc"
	"GrpcStudy/protobuf/jaeger/proto"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
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

	// jaeger 使用
	//先生成一个tracer
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{ // 采样,可以全部发送,也可以部分发送
			Type:  jaeger.SamplerTypeConst, // const 0就代表不采样,1就代表一直采样
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{ // 将这些信息放到哪一个服务器上去
			LogSpans:           true,                 // 发送span到服务器是否打印日志
			LocalAgentHostPort: "43.143.172.37:5775", // UDP端口
			//CollectorEndpoint: "http://43.143.172.37:14268/api/traces", // jaeger server
		},
		ServiceName: "mall-test",
	}

	// 基于这个configuration生成一个Trace
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger)) // 生成Trace,并且用标准输出打印日志
	if err != nil {
		panic(err)
	}
	// 将我们的这个tracer设置为全局
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	//var opts []grpc.DialOption
	//opts = append(opts, grpc.WithInsecure())
	var opts []grpc.DialOption
	//这个请求应该多长时间超时,这个重试应该几次,当服务器返回什么状态码的时候重试
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)), grpc.WithInsecure())

	//接入jaeger
	opts = append(opts, grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))

	dial, err := grpc.Dial("dns:///ws.bd.cn:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	client := proto.NewGreeterClient(dial)

	rsp, err := client.SayHello(context.Background(), &proto.Person{
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
	time.Sleep(3e9)
}
