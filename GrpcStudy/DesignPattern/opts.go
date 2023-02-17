package main

import "fmt"

func main() {
	Dial("", WithBlock())
}

func Dial(address string, opts ...DialOption) {
	cc := &ClientConn{
		dopts: dialOptions{},
	}

	for _, opt := range opts {
		opt.apply(&cc.dopts)
	}
	fmt.Println(cc.dopts.block)
}

type ClientConn struct {
	dopts dialOptions
}

type dialOptions struct {
	block bool // 是否阻塞
}

type DialOption interface {
	apply(*dialOptions)
}

type funcDialOption struct {
	f func(*dialOptions)
}

func (fdo *funcDialOption) apply(do *dialOptions) {
	fdo.f(do)
}

func WithBlock() DialOption {
	return newFuncDialOption(func(o *dialOptions) {
		o.block = true
	})
}

func newFuncDialOption(f func(*dialOptions)) *funcDialOption {
	return &funcDialOption{
		f: f,
	}
}
