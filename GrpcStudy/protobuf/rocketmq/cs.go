package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(GenerateOrderSn(10))
	}
}

func GenerateOrderSn(userId int32) string { // 并发情况下依然会出现重复的情况
	//订单号的生成规则
	/*
		年月日时分秒+用户id+2位随机数
	*/
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(),
		userId, rand.Intn(90)+10,
	)
	return orderSn
}
