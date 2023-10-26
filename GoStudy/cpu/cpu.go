package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
	"math"
	"math/rand"
	"time"
)

var globalWindow float64
var slidingMean float64

func main() {
	done := make(chan struct{})

	go func() {
		var times int
		for times <= 2 {
			globalWindow = 0
			for i := 0; i < 5; i++ {
				// 获取 CPU 利用率 (每 100 毫秒获取一次)
				percent, err := cpu.Percent(100*time.Millisecond, false)
				if err != nil {
					log.Fatalf("get CPU usage: %v\n", err)
					return
				}

				for _, v := range percent {
					fmt.Printf("every one CPU usage is %.2f%%\n", v)
					//	fmt.Println(percent)
				}
				globalWindow += percent[0]
			}

			// 计算cpu滑动均值
			if slidingMean == 0 {
				slidingMean = globalWindow / 5
			} else {
				slidingMean = (globalWindow + slidingMean) / 6
			}
			fmt.Printf("CPU usage is %.2f%%\n", slidingMean)

			times++
		}

		// 模拟程序结束后通过 channel 发送通知
		done <- struct{}{}
	}()

	// 模拟一些 CPU 密集型任务
	rand.Seed(time.Now().UnixNano())
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 10000000; i++ {
		_ = math.Sqrt(rand.Float64())
		if i == (9999999) {
			fmt.Println("end")
			break
		}
	}

	<-done
	close(done)
}
