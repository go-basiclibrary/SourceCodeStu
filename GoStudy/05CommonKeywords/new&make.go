package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//i := new(*int)
	//fmt.Println(*i)

	//md := make(map[string]string)
	//md["wangShao"] = "1"
	//go func() {
	//	time.Sleep(1e9)
	//	fmt.Println(md["wangShao"])
	//}()
	//
	//// COW
	//cmd := make(map[string]string)
	//for k, v := range md {
	//	cmd[k] = v
	//}
	//cmd["wangShao"] = "2"
	//cmd["hehe"] = "3"
	//fmt.Println(cmd)
	//fmt.Println(md)
	//
	//time.Sleep(5e9)

	//ctx := context.TODO()
	//var i int
	//fmt.Printf("1 %p,%p\n", ctx, &i)
	//TT(ctx, i)
	//fmt.Printf("3 %p,%p\n", ctx, &i)
}

//func TT(ctx context.Context, i int) {
//	fmt.Printf("2 %p,%p\n", ctx, &i)
//}

func one(ctx context.Context) {
	cancel, cancelFunc := context.WithCancel(ctx)
	cancelFunc()
	two(cancel, cancelFunc)
	fmt.Println(ctx)
}

func two(ctx context.Context, cancelFunc context.CancelFunc) {
	time.Sleep(2e9)
	fmt.Println(ctx)
}
