package main

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"math/rand"
	"time"
)

func main() {
	// 初始化sentinel
	err := api.InitDefault()
	if err != nil {
		panic(err)
	}

	// 配置限流规则
	resource := "some-test"
	_, err = flow.LoadRules([]*flow.Rule{ // 可以对规则进行多配置
		{
			Resource: resource,
			// 根据阈值进行防控,暴力
			// 系统长期处于低水平的情况下,当流量突然增加,直接把系统拉到高水位可能瞬间击垮系统
			// WarmUp 通过冷启动,让通过的流量缓慢增加,在一定时间内逐渐增加到阈值上限,避免系统被直接压垮
			TokenCalculateStrategy: flow.Direct,
			// Throttling匀速通过基于时间间隔去通过的,如果在特定间隔时间内,一个请求处理完成后,间隔内则不再处理其他请求
			ControlBehavior:  flow.Reject, // Reject超过阈值直接拒绝,Throttling匀速排队
			Threshold:        10,          // 阈值
			StatIntervalInMs: 1000,        // 粒度过大,比如5s内1w并发,但100ms可能就会达到1w并发,造成瞬间压力过大
		},
		{
			Resource:               resource + "1",
			TokenCalculateStrategy: flow.WarmUp,
			ControlBehavior:        flow.Reject,
			Threshold:              1000,
			// 预热时间长度,仅仅针对WarmUp配置才会生效
			WarmUpPeriodSec: 30,
		},
	})
	if err != nil {
		panic(err)
	}

	var globalTotal, passTotal, blockTotal int

	// sentinel WarmUp策略,会在每秒内,统计一次,通过了多少,总共多少,block了多少
	for i := 0; i < 100; i++ {
		go func() {
			for {
				globalTotal++
				e, b := api.Entry(resource+"1", api.WithTrafficType(base.Inbound))
				if b != nil {
					blockTotal++
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					passTotal++
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					e.Exit()
				}
			}
		}()
	}

	go func() {
		var oldTotal, oldPass, oldBlock int
		for {
			oneSecondTotal := globalTotal - oldTotal
			oldTotal = globalTotal
			oneSecondPass := passTotal - oldPass
			oldPass = passTotal
			oneSecondBlock := blockTotal - oldBlock
			oldBlock = blockTotal

			time.Sleep(time.Second)
			fmt.Printf("total:%d,pass:%d,block:%d\n", oneSecondTotal, oneSecondPass, oneSecondBlock)
		}
	}()

	//for i := 0; i < 12; i++ {
	//	e, b := api.Entry(resource, api.WithTrafficType(base.Inbound)) // Inbound入口流量控制
	//	if b != nil {
	//		// 违反流量控制约束
	//		fmt.Println("限流拉")
	//	} else {
	//		fmt.Println("check is through")
	//		e.Exit()
	//	}
	//}
	c := make(chan struct{})
	<-c
}
