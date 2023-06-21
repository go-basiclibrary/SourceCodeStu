package main

import (
	"fmt"
	"git.tencent.com/intl/intl_comm/intlexception"
	"git.tencent.com/trpc-go/trpc-go/log"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	//deferFunc()
	//deferZz()
	//deferZx()

	//fmt.Println(*deferTT())

	//deferFuncTwice()

	//deferFuncWithFor()

	deferRecover()
}

func deferRecover() {
	c := cron.New()
	_, err := c.AddFunc("@every 3s", taskFunc)
	if err != nil {
		log.Errorf("add func err: %v", err)
	}
	c.Start()
	time.Sleep(1 * time.Hour)
}

func taskFunc() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("go extra err : %v %s", err, intlexception.PanicStackError())
		}
	}()
	panic("123")
}

func deferFuncWithFor() {
	for i := 0; i < 3; i++ {
		procLoop()
		time.Sleep(1e9)
	}
}

func procLoop() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	proc()
}

func proc() {
	panic("proc")
}

func deferFuncTwice() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	testFunc()
}

func testFunc() {
	panic("test")

	//t := time.Now()
	//defer func() {
	//	fmt.Println(time.Since(t))
	//}()
	//defer fmt.Println(3)
	//time.Sleep(1e9)

	deferProc(3)
}

func deferProc(twice int) {
	for i := 0; i < twice; i++ {
		run()
		time.Sleep(1e9)
	}
}

func run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	proc()
}

func deferTT() *int {
	t := 0
	defer func(i *int) { *i++ }(&t)
	return &t
}

// defer 执行顺序
func deferZx() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

type N struct {
	a int
}

// 预计算指针是否会有效果
// 是会影响值的
func deferZz() {
	var s = &N{a: 100}
	defer fmt.Println(s)
	s.a = 101
}

// defer 会进行预计算
func deferFunc() {
	startAt := time.Now()
	//defer fmt.Println(time.Since(startAt)) // 预期5s但是这里只会得到0s

	// 改进方法
	defer func() {
		fmt.Println(time.Since(startAt))
	}()
	time.Sleep(5e9)
}
