package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// 创建 Cron 调度器
	c := cron.New()

	// 添加定时任务
	// 每月10日00:00定时生成计费填报
	entry1, err := c.AddFunc("0 0 10 * ?", func() {
		// 在这里执行生成计费填报的逻辑
		fmt.Println("生成计费填报")
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	// 每月15-20日10点定时通知人填报数据
	entry2, err := c.AddFunc("0 10 15-20 * ?", func() {
		// 在这里执行通知人填报数据的逻辑
		fmt.Println("通知人填报数据")
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	// 每月21日12:00定时自动提交业务
	entry3, err := c.AddFunc("0 12 21 * ?", func() {
		// 在这里执行自动提交业务的逻辑
		fmt.Println("自动提交业务")
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	// 每月27日12:00定时刷新当月计费详情所有业务状态
	entry4, err := c.AddFunc("0 12 27 * ?", func() {
		// 在这里执行刷新当月计费详情所有业务状态的逻辑
		fmt.Println("刷新当月计费详情所有业务状态")
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	// 启动 Cron 调度器
	c.Start()

	// 计算每个定时任务下一次触发的时间
	//now := time.Now()
	nextTime1 := c.Entry(entry1).Next
	nextTime2 := c.Entry(entry2).Next
	nextTime3 := c.Entry(entry3).Next
	nextTime4 := c.Entry(entry4).Next

	// 输出下一次触发的时间
	fmt.Println("下一次生成计费填报的时间:", nextTime1)
	fmt.Println("下一次通知人填报数据的时间:", nextTime2)
	fmt.Println("下一次自动提交业务的时间:", nextTime3)
	fmt.Println("下一次刷新当月计费详情所有业务状态的时间:", nextTime4)

	// 休眠，以便观察定时任务执行
	time.Sleep(5 * time.Minute)

	// 停止 Cron 调度器
	c.Stop()
}
