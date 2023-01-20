package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"gobasic/protobuf/grpc_resolver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

const (
	address     = "consul:///Test"
	defaultName = "----this is consul resolver test----"
)
const (
	timestampFormat = time.StampNano
	streamingCount  = 0
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go request(wg, int8(i))
	}

	wg.Wait()
	fmt.Println("-->end---")
}

func request(wg *sync.WaitGroup, index int8) {
	//创建client conn,以及服务器地址address
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 根据client conn,创建greeter client
	c := proto.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	for i := 0; i < 1; i++ {
		ms, err := c.SayHello(ctx, &proto.Person{Email: name})
		if err != nil {
			panic(err)
		}
		fmt.Printf("get msg %v\n", ms)
	}

	wg.Done()
}

func readTest(path string) string {
	p, _ := ioutil.ReadFile(path)
	return string(p)
}

// 重点来了,注册consul解析器
const consulResolverScheme = "consul"

// 要实现resolver builder接口
type consulResolverBuilder struct {
}

// Build 实现 builder 接口
func (*consulResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &consulResolver{
		target: target,
		cc:     cc,
	}

	// 真正的发起 链接
	r.start()

	return r, nil
}
func (*consulResolverBuilder) Scheme() string { return consulResolverScheme }

type consulResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}

// start获取后端服务列表
func (r consulResolver) start() {
	//创建consul客户端
	cfg := api.DefaultConfig()
	//配置consul服务地址
	cfg.Address = "43.143.172.37:8500"
	cli, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 服务发现
	service, _, err := cli.Health().Service("Test", "consul", true, &api.QueryOptions{})
	if err != nil {
		panic(err)
	}

	// 发现的services列表,介入到resolver.Address
	addrs := make([]resolver.Address, len(service))
	for i, s := range service {
		addr := fmt.Sprintf("%s:%d", s.Service.Address, s.Service.Port)
		addrs[i] = resolver.Address{Addr: addr}
	}

	// 更新解析器State,最终向grpc服务器端发起链接
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

// ResolveNow 实现 resolver 接口
func (*consulResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*consulResolver) Close()                                  {}

// 注册我们自己实现的解析器
func init() {
	resolver.Register(&consulResolverBuilder{})
}
