package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	//fmt.Println(time.Now())
	//local, err := time.LoadLocation("Asia/Shanghai")
	//in := time.Now().In(local)
	//fmt.Println(in)
	//go func() {
	//	c := cron.New()
	//	entry, err := c.AddFunc("18 * * * ?", func() {
	//		fmt.Println("123456")
	//	})
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	c.Start()
	//	fmt.Println(c.Entry(entry).Next)
	//	fmt.Println("cron tasks is start...")
	//}()
	//var s = []string{"46 * * * ?", "47 * * * ?", "48 * * * ?"}
	//c := cron.New()
	//for i := 0; i < 3; i++ {
	//	var t = i
	//	_, err := c.AddFunc(s[t], taskManager(s[t]))
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//c.Start()
	//
	//time.Sleep(1 * time.Hour)

	c := cron.New()
	_, err := c.AddFunc("@every 2s", func() {
		fmt.Println("abc")
	})
	if err != nil {
		panic(err)
	}
	c.Start()
	time.Sleep(100e9)
}

func taskManager(task string) func() {
	return func() {
		fmt.Println(task)
	}
}
