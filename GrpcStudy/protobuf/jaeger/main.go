package main

import (
	"github.com/opentracing/opentracing-go"
	"time"

	"github.com/uber/jaeger-client-go"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
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
	defer closer.Close()

	span := tracer.StartSpan("go-grpc-web")
	span.Finish()

	time.Sleep(1e9)
	span2 := tracer.StartSpan("testA", opentracing.ChildOf(span.Context()))
	span2.Finish() // 子span
}
