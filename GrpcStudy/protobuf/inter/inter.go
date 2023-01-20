package main

import "fmt"

// DialOptions 通过interface反向代理dailOptions,将改变内部配置的方式给暴露出来
type DialOptions interface {
	apply(*dialOptions)
}

type dialOptions struct {
	isCool bool
}

type funcDialOption struct {
	f func(*dialOptions)
}

func (fdo *funcDialOption) apply(do *dialOptions) {
	fdo.f(do)
}

func newFuncDialOption(f func(*dialOptions)) *funcDialOption {
	return &funcDialOption{
		f: f,
	}
}

func CoolStart() DialOptions {
	return newFuncDialOption(func(o *dialOptions) {
		o.isCool = true
	})
}

// 抽象一个接口
func main() {
	// 外部在启动一个东西的时候,内部可能初始化一些自己的struct(实例)
	// 那么在初始化实例的时候,我并不想提供让你直接能操作我的东西
	// 反向代理一个配置
	dial := &dialOptions{false}
	// 通过反向代理,设置isCool为true
	CoolStart().apply(dial)
	fmt.Println(dial.isCool)
}
