package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"log"
)

// manual自己手动实现解析器
func main() {
	scheme := manual.NewBuilderWithScheme("wangShao") // 构建自定义解析器
	resolver.Register(scheme)                         // 注册 scheme
	defer scheme.Close()
	// 这么做就是为了用户在客户端不需要自己手动调用
	// 而是grpc框架在build方法内部,帮我们调用UpdateState,从而触发平衡器的操作流程,
	// 即最终向grpc服务器端发起rpc链接;
	scheme.InitialState(resolver.State{ //初始化工作,内部将resolver.State赋值给r.bootstrapState
		Addresses: []resolver.Address{ //直接将后端服务器地址列表封装到resolver.State里面
			{Addr: "localhost:50051"},
			{Addr: "localhost:50052"},
		},
	})
	addr := fmt.Sprintf("%s:///unused", scheme.Scheme()) //unused,并没有使用这个进行解析后端服务地址
	opts := []grpc.DialOption{                           // DialOption 参数
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig("{}"),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("cannot dial %v", err.Error())
	}
	defer conn.Close()

}
