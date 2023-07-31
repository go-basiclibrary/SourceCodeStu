package main

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"time"
)

// errgroup使用
func main() {
	var a *messageA
	var b *messageB
	var c *messageC

	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		a = &messageA{iRet: 200, sMsg: "success"}
		return nil
	})
	group.Go(func() error {
		b = &messageB{iRet: 200, sMsg: "success"}
		return nil
	})
	group.Go(func() error {
		c = &messageC{iRet: 200, sMsg: "success"}
		return nil
	})
	err := group.Wait()
	if err != nil {
		panic(err)
	}

	// BFF 聚合
	msg := &all{
		messageA: a,
		messageB: b,
		messageC: c,
	}

	fmt.Printf("%+v\n", msg.messageA)
}

type all struct {
	*messageA
	*messageB
	*messageC
}

type messageA struct {
	iRet int
	sMsg string
}

type messageB struct {
	iRet int
	sMsg string
}

type messageC struct {
	iRet int
	sMsg string
}
