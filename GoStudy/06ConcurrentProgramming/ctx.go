package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/tal-tech/go-zero/core/contextx"
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

	// 第一版context超时实现
	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			ctx, _ := context.WithTimeout(context.Background(), 1e9)
			requestWork(ctx, "any")
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	// 查看Goroutines泄露问题  总的来说,下游的服务没有被中断,只是上游结束了
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, f := contextx.ShrinkDeadline(ctx, 2e9)
	defer f()

	done := make(chan error, 1)
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case p := <-panicChan:
		fmt.Println(p)
		return errors.New(p.(string))
	case <-ctx.Done():
		return ctx.Err()
	}
}

func hardWork(job interface{}) error {
	time.Sleep(time.Minute)
	return nil
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
