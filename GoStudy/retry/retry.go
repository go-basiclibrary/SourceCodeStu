package main

import (
	"git.tencent.com/trpc-go/trpc-go/log"
	"github.com/avast/retry-go"
	"time"
)

func main() {
	err := retry.Do(func() error {
		// 需要重试的taskFunc
		err := taskFunc()
		if err != nil {
			log.Errorf("task func err: %v", err)
			return err
		}
		return nil
	}, retry.DelayType(func(n uint, err error, config *retry.Config) time.Duration {
		// 使用指数退避算法计算重试延迟
		return time.Duration(1<<n) * time.Second
	}), retry.Attempts(3)) // 最大重试次数3

	if err != nil {
		// TODO 做一个上报
		log.Errorf("The task has been executed three times but still fails,"+
			"the err is %v", err)
	}
}

func taskFunc() error {
	return nil
}
