package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 取消上下文
	//ctx, cancel := context.WithCancel(context.TODO())
	//cancel()
	//one(ctx)
	//fmt.Println(ctx.Err())

	// 超时时,后续代码还会执行么
	//go toTimeOut()
	//select {}

	//ctx, _ := context.WithTimeout(context.TODO(), 1e9)
	//
	//go TCancelWithCtx(ctx)
	//
	//time.Sleep(2e9)
	//fmt.Println(ctx.Err())

	ctx, cancel := CtxTimeOutWithCancel(context.TODO())

	cancel()

	fmt.Println(ctx.Err())
}

func CtxTimeOutWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)

	ctx, _ = context.WithTimeout(ctx, 1e9)

	return ctx, cancel
}

func TCancelWithCtx(ctx context.Context) {
	fmt.Println("TCancelWithCtx")
	time.Sleep(1e9)
	fmt.Println("TEND")
}

func toTimeOut() {
	ctx, _ := context.WithTimeout(context.TODO(), 1e9)
	nextGoroutine(ctx)
}

func nextGoroutine(ctx context.Context) {
	nextGoroutineTwo(ctx)
	nextGoroutineTwo(ctx)
}

func nextGoroutineTwo(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Done")
	case <-time.After(2e9):

	}
}

func one(ctx context.Context) {
	cancel, cancelFunc := context.WithCancel(ctx)
	two(cancel, cancelFunc)
}

func two(ctx context.Context, cancelFunc context.CancelFunc) {
	fmt.Println(ctx.Err())
}
