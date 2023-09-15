package main

import (
	"context"
	"git.tencent.com/trpc-go/trpc-go/log"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	ctx := context.Background()

	// r:每秒可以向token桶中产生多少token
	// b:桶的容量大小
	limiter := rate.NewLimiter(10, 1)
	// r:每100ms往桶中放一个token
	//rate.NewLimiter(rate.Every(100*time.Millisecond), 1)

	// wait方法消耗token,如果拿不到token,则会等待,ctx可以决定等待最长时长
	err := limiter.Wait(ctx)
	if err != nil {
		log.Errorf("limiter wait err is %q", err.Error())
		return
	}

	// Allow
	// 截至到某一时刻,目前桶中是否至少为n个,满足则返回true,同时从桶中消费n个token
	b := limiter.AllowN(time.Now(), 1)
	if !b {
		log.Errorf("当前令牌桶不满足条件")
		return
	}

	// Reserve
	r := limiter.ReserveN(time.Now(), 1)
	if !r.OK() {
		// 如果不愿意等待,则直接退出
		log.Info("当前用户不等待")
		return
	}
	// 等待
	time.Sleep(r.Delay())

	//调整速率
	//limiter.SetLimit()
}
